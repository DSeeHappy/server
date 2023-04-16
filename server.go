package main

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/internal/infrastructure/repositories"
	"server/internal/model"
	"server/internal/service"
	"server/internal/transport/handler"
)

const (
	port = ":50051"
)

func main() {
	ctx := context.Background()
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "password"

	db, _ := newMockDB(ctx, uri, username, password)

	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	productRepo := repositories.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(*productService)

	// Create a gRPC server object
	s := grpc.NewServer()

	model.RegisterProductServiceServer(s, productHandler)

	// Serve gRPC Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func newMockDB(ctx context.Context, uri, username, password string) (neo4j.DriverWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return driver, err
	}
	defer driver.Close(ctx)
	return driver, nil
}
