package rpc

import (
	"context"

	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/services/Auth/dto"
	"github.com/RhnAdi/elearning-microservice/services/Auth/usecase"
)

type authRpc struct {
	authUsecase usecase.AuthUsecase
	*pb.UnimplementedAuthServiceServer
}

func NewAuthRpc(authUsecase usecase.AuthUsecase) pb.AuthServiceServer {
	return &authRpc{
		authUsecase: authUsecase,
	}
}

func (d *authRpc) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	convertReq := new(dto.AuthRegisterRequest)
	convertReq.FromProtoBuffer(req)

	res, err := d.authUsecase.Register(convertReq)

	return res.ToProtoBuffer(), err
}

func (d *authRpc) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	convertReq := new(dto.AuthLoginRequest)
	convertReq.FromProtoBuffer(req)

	res, err := d.authUsecase.Login(convertReq)

	return res.ToProtoBuffer(), err
}
