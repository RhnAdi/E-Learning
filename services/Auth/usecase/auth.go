package usecase

import (
	"github.com/RhnAdi/elearning-microservice/config/jwt"
	"github.com/RhnAdi/elearning-microservice/security"
	"github.com/RhnAdi/elearning-microservice/services/Auth/dto"
	"github.com/RhnAdi/elearning-microservice/services/Auth/models"
	"github.com/RhnAdi/elearning-microservice/services/Auth/repository"
	"github.com/RhnAdi/elearning-microservice/services/Auth/validators"
)

type AuthUsecase interface {
	Register(req *dto.AuthRegisterRequest) (*dto.AuthResponse, error)
	Login(req *dto.AuthLoginRequest) (*dto.AuthResponse, error)
}

type authUsecase struct {
	jwtConfig      *jwt.JWTConfig
	userRepository repository.UserRepository
}

func NewAuthUsecase(jwtConfig *jwt.JWTConfig, userRepository repository.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
		jwtConfig:      jwtConfig,
	}
}

func (u *authUsecase) Register(req *dto.AuthRegisterRequest) (*dto.AuthResponse, error) {
	if err := validators.ValidateRegisterRequest(req); err != nil {
		return &dto.AuthResponse{}, err
	}
	if _, err := u.userRepository.GetByEmail(req.Email); err == nil {
		return &dto.AuthResponse{}, validators.ErrEmailAlreadyExist
	}

	newUser := new(models.User)
	newUser.Username = req.Username
	newUser.Email = req.Email
	newUser.Password = req.Password
	newUser.Role = req.Role
	err := newUser.PreSave()
	if err != nil {
		return &dto.AuthResponse{}, err
	}

	newUser, err = u.userRepository.Save(newUser)
	if err != nil {
		return &dto.AuthResponse{}, err
	}

	data := jwt.JWTClaims{
		Id:   newUser.Id.Hex(),
		Role: newUser.Role,
	}

	access_token, err := security.CreateToken(u.jwtConfig.AccessTokenExpiresIn, data, u.jwtConfig.AccessTokenPrivateKey)
	if err != nil {
		return &dto.AuthResponse{}, err
	}
	refresh_token, err := security.CreateToken(u.jwtConfig.RefreshTokenExpiresIn, data, u.jwtConfig.RefreshTokenPrivateKey)
	if err != nil {
		return &dto.AuthResponse{}, err
	}

	res := new(dto.AuthResponse)
	res.AccessToken = access_token
	res.RefreshToken = refresh_token
	return res, err
}

func (u *authUsecase) Login(req *dto.AuthLoginRequest) (*dto.AuthResponse, error) {
	if err := validators.ValidateLoginRequest(req); err != nil {
		return &dto.AuthResponse{}, err
	}

	user, err := u.userRepository.GetByEmail(req.Email)
	if err != nil {
		return &dto.AuthResponse{}, validators.ErrUserNotFound
	}

	err = user.VerifyPassword(req.Password)
	if err != nil {
		return &dto.AuthResponse{}, err
	}

	data := jwt.JWTClaims{
		Id:   user.Id.Hex(),
		Role: user.Role,
	}

	access_token, err := security.CreateToken(u.jwtConfig.AccessTokenExpiresIn, data, u.jwtConfig.AccessTokenPrivateKey)
	if err != nil {
		return &dto.AuthResponse{}, err
	}
	refresh_token, err := security.CreateToken(u.jwtConfig.RefreshTokenExpiresIn, data, u.jwtConfig.RefreshTokenPrivateKey)
	if err != nil {
		return &dto.AuthResponse{}, err
	}

	res := new(dto.AuthResponse)
	res.AccessToken = access_token
	res.RefreshToken = refresh_token
	return res, err
}
