package service

import (
	"context"
	"errors"
	"server/internal/infrastructure/repositories"
	"server/internal/model"
)

// Business Logic - Product

type ProductService struct {
	model.UnsafeProductServiceServer
	repo productService
}

type productService interface {
	repositories.ProductRepository
}

func (p ProductService) CreateProduct(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {
	// Validate the product fields
	if request.Product.Name == "" {
		return nil, errors.New("invalid product data")
	}
	product := &model.Product{
		Name:  request.Product.Name,
		Price: request.Product.Price,
	}
	// save to db
	productResponse, err := p.repo.CreateProduct(ctx, &model.ProductCreateRequest{Product: product})
	if err != nil {
		return nil, err
	}
	return productResponse, nil
}

func (p ProductService) ReadProduct(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) UpdateProduct(ctx context.Context, request *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) DeleteProduct(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductService) ListProducts(ctx context.Context, product *model.Product) (*model.ProductListResponse, error) {
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
