package repositories

import (
	"github.com/kryast/crud-14.git/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (pr *productRepository) Create(product *models.Product) error {
	return pr.db.Create(product).Error
}

func (pr *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := pr.db.Find(&products).Error

	return products, err

}

func (pr *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product

	err := pr.db.First(&product, id).Error

	return &product, err
}

func (pr *productRepository) Update(product *models.Product) error {
	return pr.db.Save(product).Error
}
