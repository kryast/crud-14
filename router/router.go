package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-14.git/handlers"
	"github.com/kryast/crud-14.git/repositories"
	"github.com/kryast/crud-14.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewProductRepository(db)
	svc := services.NewProductService(repo)
	h := handlers.NewProductHandler(svc)

	r.PUT("/product", h.Create)
	r.GET("/product", h.GetAll)
	r.GET("/product/:id", h.GetByID)
	r.PUT("/product/:id", h.Update)
	r.DELETE("/product/:id", h.Delete)

	return r

}
