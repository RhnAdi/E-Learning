package repository_test

import (
	"log"
	"testing"
	"time"

	db "github.com/RhnAdi/elearning-microservice/config/database"
	"github.com/RhnAdi/elearning-microservice/services/Auth/models"
	"github.com/RhnAdi/elearning-microservice/services/Auth/repository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln("cant load env var : ", err.Error())
	}
}

func TestUserRespositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	user := models.User{
		Username: "test_username",
		Email:    "test@email",
		Password: "test",
		Role:     "student",
	}
	err = user.PreSave()
	assert.NoError(t, err)
	createUser, err := userRepository.Save(&user)
	assert.NoError(t, err)
	assert.NotNil(t, createUser)
	assert.Equal(t, createUser.Username, user.Username)
	assert.Equal(t, createUser.Email, user.Email)
	assert.Equal(t, createUser.Password, user.Password)
	assert.Equal(t, createUser.Role, user.Role)
	assert.Equal(t, createUser.Role, user.Role)
}

func TestUserRepositoryGetById(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)

	user := models.User{
		Username: "test_get_by_id_username",
		Email:    "test_get_by_id@email",
		Password: "test_get_by_id",
		Role:     "student",
	}
	err = user.PreSave()
	assert.NoError(t, err)

	_, err = userRepository.Save(&user)
	assert.NoError(t, err)

	getUser, err := userRepository.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.NotNil(t, getUser)
	assert.Equal(t, getUser.Username, user.Username)
	assert.Equal(t, getUser.Email, user.Email)
	assert.Equal(t, getUser.Password, user.Password)
	assert.Equal(t, getUser.Role, user.Role)
}

func TestUserReporsitoryGetByEmail(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)

	user := models.User{
		Username: "test_get_by_email_username",
		Email:    "test_get_by_email@email",
		Password: "test_get_by_email",
		Role:     "student",
	}
	err = user.PreSave()
	assert.NoError(t, err)

	_, err = userRepository.Save(&user)
	assert.NoError(t, err)

	getUser, err := userRepository.GetByEmail(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, getUser)
	assert.Equal(t, getUser.Username, user.Username)
	assert.Equal(t, getUser.Email, user.Email)
	assert.Equal(t, getUser.Password, user.Password)
	assert.Equal(t, getUser.Role, user.Role)
}

func TestUserReporsitoryUpdate(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)

	user := models.User{
		Username: "test_get_by_email_username",
		Email:    "test_get_by_email@email",
		Password: "test_get_by_email",
		Role:     "student",
	}
	err = user.PreSave()
	assert.NoError(t, err)

	_, err = userRepository.Save(&user)
	assert.NoError(t, err)

	user.Username = "test_update_username"
	user.Email = "test_update_email@email"
	user.Password = "test_update_password"
	user.Role = "teacher"
	user.UpdatedAt = time.Now()

	updateUser, err := userRepository.Update(user)
	assert.NoError(t, err)
	assert.NotNil(t, updateUser)
	assert.Equal(t, updateUser.Username, user.Username)
	assert.Equal(t, updateUser.Email, user.Email)
	assert.Equal(t, updateUser.Password, user.Password)
	assert.Equal(t, updateUser.Role, user.Role)
	assert.Equal(t, updateUser.UpdatedAt, user.UpdatedAt)
}

func TestUserRepositoryDelete(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)

	user := models.User{
		Username: "test_delete_username",
		Email:    "test_delete@email",
		Password: "test_delete",
		Role:     "student",
	}
	err = user.PreSave()
	assert.NoError(t, err)

	_, err = userRepository.Save(&user)
	assert.NoError(t, err)

	deleteUser, err := userRepository.Delete(user)
	assert.NoError(t, err)
	assert.NotNil(t, deleteUser)
	assert.Equal(t, deleteUser.Username, user.Username)
	assert.Equal(t, deleteUser.Email, user.Email)
	assert.Equal(t, deleteUser.Password, user.Password)
	assert.Equal(t, deleteUser.Role, user.Role)
	assert.Equal(t, deleteUser.UpdatedAt, user.UpdatedAt)
}

func TestUserRespositoryGetAll(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)

	_, err = userRepository.GetAll()

	assert.NoError(t, err)
}
