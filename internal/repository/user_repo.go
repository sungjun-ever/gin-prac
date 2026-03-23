package repository

import (
	"gorm.io/gorm"
	"play/internal/model"
)

// UserRepo가 구현해야할 메서드 선언
type UserRepo interface {
	Create(user *model.User) error
	FindById(id uint) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}
