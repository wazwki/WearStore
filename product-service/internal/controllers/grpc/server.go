package server

import (
	"github.com/wazwki/WearStore/product-service/api/proto/productpb"
	"github.com/wazwki/WearStore/product-service/internal/services"
)

type Server struct {
	productpb.UnimplementedProductServiceServer
	service services.ServiceInterface
}

func NewServer(s services.ServiceInterface) productpb.ProductServiceServer {
	return &Server{service: s}
}
