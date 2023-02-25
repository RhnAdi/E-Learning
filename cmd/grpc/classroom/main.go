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
	"github.com/RhnAdi/elearning-microservice/services/Classroom/delivery/rpc"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/repository"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/usecase"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(fmt.Errorf("(classroom) unable to load env : %w", err))
	}

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	dbConfig := db.NewConfig()
	conn, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Panicln(fmt.Errorf("(classroom) can't connect to database : %w", err))
	}

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)
	classroomRpc := rpc.NewClassroomRpc(classroomUsecase)

	dns := fmt.Sprintf("%s:%s", appHost, appPort)
	lis, err := net.Listen("tcp", dns)
	if err != nil {
		log.Panicln(fmt.Errorf("(classroom) can't listen tcp : %s : %w", dns, err))
	}
	classroomServicePath := "/Classroom.ClassroomService/"
	accessibleRole := map[string][]string{
		classroomServicePath + "CreateClassroom":   {"teacher"},
		classroomServicePath + "UpdateClassroom":   {"teacher"},
		classroomServicePath + "DeleteClassroom":   {"teacher"},
		classroomServicePath + "GetAllJoinRequest": {"teacher"},
		classroomServicePath + "AcceptJoinRequest": {"teacher"},
		classroomServicePath + "RejectJoinRequest": {"teacher"},
		classroomServicePath + "JoinClass":         {"student"},
		classroomServicePath + "MyClass":           {"teacher", "student"},
		classroomServicePath + "GetStudentInfo":    {"teacher", "student"},
	}
	jwtconfig, err := jwt.NewJWTConfig("../../../")
	if err != nil {
		log.Panicln(fmt.Errorf("(classroom) unable to create jwt config"))
	}
	authInterceptor := interceptor.NewAuthInterceptor(&jwtconfig, accessibleRole)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.Unary()))
	pb.RegisterClassroomServiceServer(grpcServer, classroomRpc)
	log.Println("(classroom) service server running on :", dns)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panicln(fmt.Errorf("(classroom) can't running grpc server : %w", err))
	}
}
