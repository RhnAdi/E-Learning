package usecase_test

import (
	"log"
	"testing"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/config/jwt"
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

func TestRegisterUserUsecase(t *testing.T) {
	dbcfg := db.NewConfig()
	authcfg, err := jwt.NewJWTConfig()
	assert.NoError(t, err)

	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	authUsecase := usecase.NewAuthUsecase(&authcfg, userRepository)

	req := dto.AuthRegisterRequest{
		Username: "admin",
		Email:    "admin@gmail.com",
		Password: "12345678",
		Role:     "admin",
	}

	_, err = authUsecase.Register(&req)
	assert.NoError(t, err)
}

func TestLoginUserUsecase(t *testing.T) {
	dbcfg := db.NewConfig()
	authcfg, err := jwt.NewJWTConfig()
	assert.NoError(t, err)

	conn, err := db.NewConnection(dbcfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	authUsecase := usecase.NewAuthUsecase(&authcfg, userRepository)

	regReq := dto.AuthRegisterRequest{
		Username: "test_login",
		Email:    "test_login@email",
		Password: "test_password",
		Role:     "teacher",
	}

	_, err = authUsecase.Register(&regReq)
	assert.NoError(t, err)

	loginRequest := dto.AuthLoginRequest{
		Email:    "test_login@email",
		Password: "test_password",
	}

	_, err = authUsecase.Login(&loginRequest)
	assert.NoError(t, err)
}
