package usecase

import (
	"github.com/RhnAdi/elearning-microservice/services/Auth/dto"
	"github.com/RhnAdi/elearning-microservice/services/Auth/models"
	"github.com/RhnAdi/elearning-microservice/services/Auth/repository"
	"github.com/RhnAdi/elearning-microservice/services/Auth/validators"
)

type UserUsecase interface {
	CreateUser(*dto.CreateUserRequest) (*dto.UserResponse, error)
	GetUserById(id string) (*dto.UserResponse, error)
	GetUserByEmail(email string) (*dto.UserResponse, error)
	GetAllUser() (*dto.ListUserResponse, error)
	UpdateUser(id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	DeleteUser(id string) (*dto.UserResponse, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func (u *userUsecase) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	if err := validators.ValidateCreateUserRequest(req); err != nil {
		return &dto.UserResponse{}, err
	}
	if _, err := u.userRepository.GetByEmail(req.Email); err == nil {
		return &dto.UserResponse{}, validators.ErrEmailAlreadyExist
	}
	newUser := new(models.User)
	newUser.Username = req.Username
	newUser.Email = req.Email
	newUser.Password = req.Password
	newUser.Role = req.Role
	err := newUser.PreSave()
	if err != nil {
		return &dto.UserResponse{}, err
	}

	user, err := u.userRepository.Save(newUser)

	res := new(dto.UserResponse)
	res.FromUserModel(*user)
	return res, err
}

func (u *userUsecase) GetUserById(id string) (*dto.UserResponse, error) {
	if err := validators.ValidateId(id); err != nil {
		return &dto.UserResponse{}, err
	}

	user, err := u.userRepository.GetById(id)

	res := new(dto.UserResponse)
	res.FromUserModel(user)
	return res, err
}

func (u *userUsecase) GetAllUser() (*dto.ListUserResponse, error) {
	users, err := u.userRepository.GetAll()

	res := new(dto.ListUserResponse)
	for _, user := range users {
		temp := new(dto.UserResponse)
		temp.FromUserModel(user)
		res.Users = append(res.Users, temp)
	}
	return res, err
}

func (u *userUsecase) GetUserByEmail(email string) (*dto.UserResponse, error) {
	user, err := u.userRepository.GetByEmail(email)

	res := new(dto.UserResponse)
	res.FromUserModel(user)
	return res, err
}

func (u *userUsecase) UpdateUser(id string, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	if err := validators.ValidateUpdateUserRequest(id, req); err != nil {
		return &dto.UserResponse{}, err
	}

	user, err := u.userRepository.GetById(id)
	if err != nil {
		return &dto.UserResponse{}, validators.ErrUserNotFound
	}

	user.Username = req.Username

	user, err = u.userRepository.Update(user)

	res := new(dto.UserResponse)
	res.FromUserModel(user)
	return res, err
}

func (u *userUsecase) DeleteUser(id string) (*dto.UserResponse, error) {
	if err := validators.ValidateId(id); err != nil {
		return &dto.UserResponse{}, err
	}

	user, err := u.userRepository.GetById(id)
	if err != nil {
		return &dto.UserResponse{}, validators.ErrUserNotFound
	}

	user, err = u.userRepository.Delete(user)

	res := new(dto.UserResponse)
	res.FromUserModel(user)
	return res, err
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}

}
