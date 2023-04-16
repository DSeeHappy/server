package main

import (
	"github.com/dseehappy/server/product/productpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) CreateProduct(ctx context.Context, in *CreateProductRequest) (*CreateProductResponse, error) {
	panic("implement me")
}

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Ping service to the server
	model.product.pb.RegisterProductServiceServer(s, &server{})
	// Serve gRPC Server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
