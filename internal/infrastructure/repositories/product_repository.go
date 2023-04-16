package repositories

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// DB Logic - Product

type ProductRepository interface {
}

type productRepository struct {
	db *neo4j.DriverWithContext
}

func NewProductRepository(db *neo4j.DriverWithContext) ProductRepository {
	return &productRepository{db: db}
}
