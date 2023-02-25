package main

import (
	"fmt"
	"log"
	"net"
	"os"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/config/jwt"
	"github.com/RhnAdi/elearning-microservice/pb"
	"github.com/RhnAdi/elearning-microservice/pkg/interceptor"
	"github.com/RhnAdi/elearning-microservice/services/Auth/delivery/rpc"
	"github.com/RhnAdi/elearning-microservice/services/Auth/repository"
	"github.com/RhnAdi/elearning-microservice/services/Auth/usecase"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(fmt.Errorf("(auth) unable to load env : %w", err))
	}

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(fmt.Errorf("(auth) can't connect to database : %w", err))
	}
	defer conn.Close()
	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userRpc := rpc.NewUserRpc(userUsecase)
	jwtCfg, err := jwt.NewJWTConfig("../../../")
	if err != nil {
		log.Panicln(fmt.Errorf("(auth) unable to create jwt config : %w", err))
	}
	authUsecase := usecase.NewAuthUsecase(&jwtCfg, userRepository)
	authRpc := rpc.NewAuthRpc(authUsecase)

	dns := fmt.Sprintf("%s:%s", appHost, appPort)
	lis, err := net.Listen("tcp", dns)
	if err != nil {
		log.Panicln(fmt.Errorf("(auth) can't run app on : %s : %w", dns, err))
	}

	userServicePath := "/User.UserService/"
	accessibleRoles := map[string][]string{
		userServicePath + "GetAllUser": {"admin"},
		userServicePath + "DeleteUser": {"admin"},
	}
	authInterceptor := interceptor.NewAuthInterceptor(&jwtCfg, accessibleRoles)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()))
	pb.RegisterUserServiceServer(grpcServer, userRpc)
	pb.RegisterAuthServiceServer(grpcServer, authRpc)
	log.Println("(auth) service server running on :", dns)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panicln(fmt.Errorf("(auth) server can't running : %w", err))
	}
}
