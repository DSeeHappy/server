package repositories

import "server/internal/model"

type ProductRepository interface {
	Create(product *model.Product) error
}

type productRepository struct {
	db *neo4j.DriverWithContext
}
