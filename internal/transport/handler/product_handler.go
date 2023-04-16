package handler

import (
	"server/internal/model"
	"server/internal/service"
)

type ProductHandler struct {
	model.UnimplementedProductServiceServer
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}
