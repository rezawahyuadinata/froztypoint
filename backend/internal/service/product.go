package service

import (
	"froztypoint/backend/internal/domain"
	"froztypoint/backend/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) List() ([]domain.Product, error) {
	// Bisa tambah validasi, caching, logging, dll di sini
	return s.repo.List()
}
