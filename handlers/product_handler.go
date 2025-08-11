package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-14.git/models"
	"github.com/kryast/crud-14.git/services"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

func (ph *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err := ph.service.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) GetAll(c *gin.Context) {

	products, err := ph.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, products)
}
