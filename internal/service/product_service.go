package service

import (
	"context"
	"server/internal/infrastructure/repositories"
	"server/internal/model"
)

// Business Logic - Product

type ProductService struct {
	model.UnimplementedProductServiceServer
	repo repositories.ProductRepository
}

func (p ProductService) Create(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) Read(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) Update(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) Delete(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) List(ctx context.Context, product *model.Product) (*model.ProductListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewProductService(repo repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
