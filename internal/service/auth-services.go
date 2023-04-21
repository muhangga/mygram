package service

import (
	"errors"

	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
	"github.com/muhangga/internal/utils"
)

type AuthService interface {
	Login(loginDTO entity.LoginDTO) (entity.Login, error)
	Register(registerDTO entity.RegisterDTO) (entity.User, error)
	IsEmailAvailable(email string) bool
}

type authService struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
}

func NewAuthService(authRepository repository.AuthRepository, userRepository repository.UserRepository) AuthService {
	return &authService{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

func (s *authService) Login(loginDTO entity.LoginDTO) (entity.Login, error) {
	var login entity.Login
	user, err := s.userRepository.FindByEmail(loginDTO.Email)
	if err != nil {
		return login, errors.New("user not found")
	}

	if verifiedPassword := utils.VerifyPassword(user.Password, loginDTO.Password); !verifiedPassword {
		return login, errors.New("invalid password")
	}

	generateToken, err := NewJwtService().GenerateToken(user.ID)
	if err != nil {
		return login, err
	}

	login.ID = user.ID
	login.Username = user.Username
	login.Email = user.Email
	login.Age = user.Age
	login.Token = generateToken

	return login, nil
}

func (s *authService) Register(registerDTO entity.RegisterDTO) (entity.User, error) {

	if registerDTO.Age <= 8 {
		return entity.User{}, errors.New("age must be greater than 8")
	}

	if len(registerDTO.Password) <= 6 {
		return entity.User{}, errors.New("password must be greater than 6")
	}

	user := entity.User{
		Email:    registerDTO.Email,
		Password: utils.HashPassword(registerDTO.Password),
		Username: registerDTO.Username,
		Age:      registerDTO.Age,
	}

	if err := s.authRepository.Save(&user); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *authService) IsEmailAvailable(email string) bool {
	isExist, _ := s.userRepository.IsEmailExist(email)
	return !isExist
}
