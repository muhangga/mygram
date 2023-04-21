package service

import (
	"errors"

	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/repository"
)

type UserService interface {
	GetUserByID(id int) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
func (s userService) GetUserByID(id int) (entity.User, error) {
	users, err := s.userRepository.FindUserById(id)
	if err != nil {
		return entity.User{}, err
	}

	if users.ID == 0 {
		return entity.User{}, errors.New("user not found")
	}

	return users, nil
}
