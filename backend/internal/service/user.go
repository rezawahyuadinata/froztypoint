package service

import (
	"froztypoint/backend/internal/repository"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetUserByID(id uint) (*repository.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(user *repository.User) error {
	return s.repo.Create(user)
}
