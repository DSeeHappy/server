package repositories

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/internal/model"
)

// DB Logic - Product

type ProductRepository interface {
	List(ctx context.Context, req *model.ProductListResponse) (*model.ProductListResponse, error)
	Create(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error)
	Read(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error)
	Update(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error)
	Delete(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error)
}

type productRepository struct {
	db *neo4j.DriverWithContext
}

func (p productRepository) List(ctx context.Context, req *model.ProductListResponse) (*model.ProductListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Create(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {

	return nil, nil
}

func (p productRepository) Read(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Update(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productRepository) Delete(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db neo4j.DriverWithContext) ProductRepository {
	return &productRepository{db: &db}
}
