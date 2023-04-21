package repository

import (
	"github.com/muhangga/internal/entity"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(user *entity.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Save(user *entity.User) error {
	return r.db.Debug().Create(user).Error
}
