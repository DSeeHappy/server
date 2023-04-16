package repositories

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"server/internal/model"
)

// DB Logic - Product

type ProductRepository interface {
	ListProducts(ctx context.Context, req *model.ProductListResponse) (*model.ProductListResponse, error)
	CreateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error)
	GetProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error)
	UpdateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error)
	DeleteProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error)
}

type productRepository struct {
	db *neo4j.DriverWithContext
}

func (r *productRepository) ListProducts(ctx context.Context, req *model.ProductListResponse) (*model.ProductListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *productRepository) CreateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error) {
	// Create a new session and begin a transaction
	session := (*r.db).NewSession(ctx, neo4j.SessionConfig{})
	defer func(session neo4j.SessionWithContext, ctx context.Context) {
		err := session.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(session, ctx)
	tx, err := session.BeginTransaction(ctx, neo4j.TransactionConfig{})
	if err != nil {
		return nil, err
	}
	defer func(tx neo4j.ExplicitTransaction, ctx context.Context) {
		err := tx.Rollback(ctx)
		if err != nil {
			panic(err)
		}
	}(tx, ctx)

	// Create a new product node in the database
	result, err := tx.Run(ctx, "CREATE (p:Product {name: $name, price: $price}) RETURN id(p)", map[string]interface{}{
		"name":  req.Product.Name,
		"price": req.Product.Price,
	})
	if err != nil {
		return nil, err
	}
	if !result.Next(ctx) {
		return nil, errors.New("failed to create product")
	}
	product := result.Record().Values()[0].(int64)

	// Commit the transaction and return the new product ID
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *productRepository) UpdateProduct(ctx context.Context, req *model.ProductCreateRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r productRepository) DeleteProduct(ctx context.Context, req *model.ProductRequest) (*model.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductRepository(db neo4j.DriverWithContext) ProductRepository {
	return &productRepository{db: &db}
}
