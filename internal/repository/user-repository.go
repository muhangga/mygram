package repository

import (
	"github.com/muhangga/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	IsEmailExist(email string) (bool, error)
	FindByEmail(email string) (entity.User, error)
	FindUserById(id int) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) IsEmailExist(email string) (bool, error) {
	var user entity.User
	err := r.db.Select("email").Where("email = ?", email).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindUserById(id int) (entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}
