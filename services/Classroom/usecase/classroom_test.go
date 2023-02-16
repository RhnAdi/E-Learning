package usecase_test

import (
	"log"
	"testing"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/dto"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/repository"
	"github.com/RhnAdi/elearning-microservice/services/Classroom/usecase"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("cant load env var : ", err.Error())
	}
}

func TestClassroomUsecaseCreate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)

	req := dto.CreateClassroomRequest{
		Name:        "Classroom Usecase Test 1",
		Description: "Description Usecase Test 1",
		TeacherId:   primitive.NewObjectID(),
	}

	res, err := classroomUsecase.CreateClassroom(&req)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, req.Name)
	assert.Equal(t, res.Description, req.Description)
	assert.Equal(t, res.TeacherId.Hex(), req.TeacherId.Hex())

}
func TestClassroomUsecaseGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)

	req := dto.CreateClassroomRequest{
		Name:        "Classroom Usecase Test Get By Id",
		Description: "Description Usecase Test Get By Id",
		TeacherId:   primitive.NewObjectID(),
	}

	res, err := classroomUsecase.CreateClassroom(&req)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, req.Name)
	assert.Equal(t, res.Description, req.Description)
	assert.Equal(t, res.TeacherId.Hex(), req.TeacherId.Hex())

	result, err := classroomUsecase.GetClassroomById(res.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, result.Name, req.Name)
	assert.Equal(t, result.Description, req.Description)
	assert.Equal(t, result.TeacherId.Hex(), req.TeacherId.Hex())
}

func TestClassroomUsecaseGetByName(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)

	req := dto.CreateClassroomRequest{
		Name:        "Classroom Usecase Test Get By Name",
		Description: "Description Usecase Test Get By Name",
		TeacherId:   primitive.NewObjectID(),
	}

	res, err := classroomUsecase.CreateClassroom(&req)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, req.Name)
	assert.Equal(t, res.Description, req.Description)
	assert.Equal(t, res.TeacherId.Hex(), req.TeacherId.Hex())

	result, err := classroomUsecase.GetClassroomByName(res.Name)
	assert.NoError(t, err)
	assert.Equal(t, result.Name, req.Name)
	assert.Equal(t, result.Description, req.Description)
	assert.Equal(t, result.TeacherId.Hex(), req.TeacherId.Hex())

}
func TestClassroomUsecaseUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)

	req := dto.CreateClassroomRequest{
		Name:        "Classroom Usecase Test Update",
		Description: "Description Usecase Test Update",
		TeacherId:   primitive.NewObjectID(),
	}

	res, err := classroomUsecase.CreateClassroom(&req)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, req.Name)
	assert.Equal(t, res.Description, req.Description)
	assert.Equal(t, res.TeacherId.Hex(), req.TeacherId.Hex())

	reqUpdate := dto.UpdateClassroomRequest{
		Name:        "Classroom Usecase Test Updated",
		Description: "Description Usecase Test Updated",
	}
	result, err := classroomUsecase.UpdateClassroom(res.Id.Hex(), &reqUpdate)
	assert.NoError(t, err)
	assert.Equal(t, result.Name, req.Name)
	assert.Equal(t, result.Description, req.Description)
	assert.Equal(t, result.TeacherId.Hex(), req.TeacherId.Hex())
}

func TestClassroomUsecaseDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)

	req := dto.CreateClassroomRequest{
		Name:        "Classroom Usecase Test Delete",
		Description: "Description Usecase Test Delete",
		TeacherId:   primitive.NewObjectID(),
	}

	res, err := classroomUsecase.CreateClassroom(&req)
	assert.NoError(t, err)
	assert.Equal(t, res.Name, req.Name)
	assert.Equal(t, res.Description, req.Description)
	assert.Equal(t, res.TeacherId.Hex(), req.TeacherId.Hex())

	result, err := classroomUsecase.DeleteClassroom(res.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, result.Name, req.Name)
	assert.Equal(t, result.Description, req.Description)
	assert.Equal(t, result.TeacherId.Hex(), req.TeacherId.Hex())

}

func TestClassroomUsecaseJoinClass(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	classroomRepository := repository.NewClassroomRepository(conn)
	classroomUsecase := usecase.NewClassroomUsecase(classroomRepository)

	classReq := dto.CreateClassroomRequest{
		Name:        "My Classroom",
		Description: "desct ex hash binary clear root playing craft mode happy join motto day",
		TeacherId:   primitive.NewObjectID(),
	}

	class, err := classroomUsecase.CreateClassroom(&classReq)
	assert.NoError(t, err)

	req := dto.JoinClassRequest{
		ClassroomId: class.Id,
		StudentId:   primitive.NewObjectID(),
	}

	res, err := classroomUsecase.JoinClass(&req)
	assert.NoError(t, err)
	assert.Equal(t, req.ClassroomId.Hex(), res.ClassroomId.Hex())
	assert.Equal(t, req.StudentId.Hex(), res.StudentId.Hex())
}
