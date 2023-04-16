package service

import (
	"server/internal/infrastructure/repositories"
)

// Business Logic - Product

type ProductService interface {
	repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}
