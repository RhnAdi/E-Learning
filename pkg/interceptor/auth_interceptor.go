package interceptor

import (
	"context"

	"github.com/RhnAdi/elearning-microservice/config/jwt"
	"github.com/RhnAdi/elearning-microservice/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type CtxKey string

type AuthInterceptor interface {
	Unary() grpc.UnaryServerInterceptor
}

type authInterceptor struct {
	JWTConfig       *jwt.JWTConfig
	AccessibleRoles map[string][]string
}

func NewAuthInterceptor(JWTConfig *jwt.JWTConfig, AccessibleRoles map[string][]string) AuthInterceptor {
	return &authInterceptor{
		JWTConfig:       JWTConfig,
		AccessibleRoles: AccessibleRoles,
	}
}

func (i *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (i *authInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	accessibleRole, ok := i.AccessibleRoles[method]
	if !ok {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Error(codes.Unauthenticated, "metadata is not provide")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Error(codes.Unauthenticated, "authorization token is not provide")
	}

	accessToken := values[0]
	sub, err := security.ValidateToken(accessToken, i.JWTConfig.AccessTokenPublicKey)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid : %w", err.Error())
	}

	claims := sub.(map[string]interface{})
	for _, role := range accessibleRole {
		if role == claims["Role"] {
			ctx = context.WithValue(ctx, CtxKey("claim"), sub)
			return ctx, nil
		}
	}

	return ctx, status.Error(codes.PermissionDenied, "no permission to access service")
}
