package repository

import (
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *User) error {
	return r.db.Create(user).Error
}
