package repository

import (
	"froztypoint/backend/internal/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int64
}

func (p *Product) ToDomain() domain.Product {
	return domain.Product{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}

type ProductRepository interface {
	List() ([]domain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) List() ([]domain.Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	domainProducts := make([]domain.Product, len(products))
	for i, p := range products {
		domainProducts[i] = p.ToDomain()
	}
	return domainProducts, nil
}
