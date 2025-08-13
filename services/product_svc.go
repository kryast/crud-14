package services

import (
	"github.com/kryast/crud-14.git/models"
	"github.com/kryast/crud-14.git/repositories"
)

type ProductService interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (ps *productService) Create(product *models.Product) error {
	return ps.repo.Create(product)
}

func (ps *productService) GetAll() ([]models.Product, error) {
	return ps.repo.GetAll()
}

func (ps *productService) GetByID(id uint) (*models.Product, error) {
	return ps.repo.GetByID(id)
}

func (ps *productService) Update(product *models.Product) error {
	return ps.repo.Update(product)
}
