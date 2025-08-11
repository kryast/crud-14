package services

import "github.com/kryast/crud-14.git/repositories"

type ProductService interface{}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}
