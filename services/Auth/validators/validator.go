package validators

import (
	"errors"

	"github.com/RhnAdi/elearning-microservice/services/Auth/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrInvalidId         = errors.New("invalid userId")
	ErrEmptyUsername     = errors.New("username can't be empty")
	ErrEmptyEmail        = errors.New("email can't be empty")
	ErrEmptyPassword     = errors.New("password can't be empty")
	ErrEmailAlreadyExist = errors.New("email already exist")
	ErrUserNotFound      = errors.New("user not found")
)

func ValidateLoginRequest(req *dto.AuthLoginRequest) error {
	if req.Email == "" {
		return ErrEmptyEmail
	}

	if req.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

func ValidateRegisterRequest(req *dto.AuthRegisterRequest) error {
	if req.Username == "" {
		return ErrEmptyUsername
	}

	if req.Email == "" {
		return ErrEmptyEmail
	}

	if req.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

func ValidateCreateUserRequest(req *dto.CreateUserRequest) error {
	if req.Username == "" {
		return ErrEmptyUsername
	}

	if req.Email == "" {
		return ErrEmptyEmail
	}

	if req.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

func ValidateId(id string) error {
	if !primitive.IsValidObjectID(id) {
		return ErrInvalidId
	}
	return nil
}

func ValidateUpdateUserRequest(id string, req *dto.UpdateUserRequest) error {
	err := ValidateId(id)
	if req.Username == "" {
		return ErrEmptyUsername
	}

	return err
}
