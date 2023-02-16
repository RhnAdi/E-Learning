package usecase_test

import (
	"log"
	"testing"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Auth/dto"
	"github.com/RhnAdi/elearning-microservice/services/Auth/repository"
	"github.com/RhnAdi/elearning-microservice/services/Auth/usecase"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("cant load env var : ", err.Error())
	}
}

func TestUserUsecaseCreate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	req := dto.CreateUserRequest{
		Username: "test_usecase_create_username",
		Email:    "test_usecase_create@email",
		Password: "test_usecase_create_password",
		Role:     "student",
	}

	res, err := userUsecase.CreateUser(&req)
	assert.NoError(t, err)
	assert.Equal(t, res.Username, req.Username)
	assert.Equal(t, res.Email, req.Email)
	assert.Equal(t, res.Role, req.Role)
}

func TestUserUsecaseGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	req := dto.CreateUserRequest{
		Username: "test_usecase_GetById_username",
		Email:    "test_usecase_GetById@email",
		Password: "test_usecase_GetById_password",
		Role:     "student",
	}

	created, err := userUsecase.CreateUser(&req)
	assert.NoError(t, err)

	res, err := userUsecase.GetUserById(created.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, res.Username, req.Username)
	assert.Equal(t, res.Email, req.Email)
	assert.Equal(t, res.Role, req.Role)
}

func TestUserUsecaseGetByEmail(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	req := dto.CreateUserRequest{
		Username: "test_usecase_GetByEmail_username",
		Email:    "test_usecase_GetByEmail@email",
		Password: "test_usecase_GetByEmail_password",
		Role:     "student",
	}

	created, err := userUsecase.CreateUser(&req)
	assert.NoError(t, err)

	res, err := userUsecase.GetUserByEmail(created.Email)
	assert.NoError(t, err)
	assert.Equal(t, res.Username, req.Username)
	assert.Equal(t, res.Email, req.Email)
	assert.Equal(t, res.Role, req.Role)
}

func TestUserUsecaseUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	req := dto.CreateUserRequest{
		Username: "test_usecase_update_username",
		Email:    "test_usecase_update@email",
		Password: "test_usecase_update_password",
		Role:     "student",
	}

	created, err := userUsecase.CreateUser(&req)
	assert.NoError(t, err)

	updateReq := dto.UpdateUserRequest{
		Username: "test_usecase_update_username(updated)",
	}
	res, err := userUsecase.UpdateUser(created.Id.Hex(), &updateReq)
	assert.NoError(t, err)
	assert.Equal(t, res.Username, updateReq.Username)
}

func TestUserUsecaseDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	req := dto.CreateUserRequest{
		Username: "test_usecase_delete_username",
		Email:    "test_usecase_delete@email",
		Password: "test_usecase_delete_password",
		Role:     "student",
	}

	created, err := userUsecase.CreateUser(&req)
	assert.NoError(t, err)

	res, err := userUsecase.DeleteUser(created.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, res.Id.Hex(), created.Id.Hex())
}
