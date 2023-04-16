package repositories

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/internal/model"
)

// DB Logic - Product

type ProductRepository interface {
	ListProducts(ctx context.Context, req *model.ProductListResponse) (*model.ProductListResponse, error)
	CreateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error)
	ReadProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error)
	UpdateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error)
	DeleteProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error)
}

type productRepository struct {
	db *neo4j.DriverWithContext
}

func (p productRepository) ListProducts(ctx context.Context, req *model.ProductListResponse) (*model.ProductListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) CreateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) ReadProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) UpdateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) DeleteProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db neo4j.DriverWithContext) ProductRepository {
	return &productRepository{db: &db}
}
