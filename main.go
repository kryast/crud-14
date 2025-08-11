package main

import (
	"log"

	"github.com/kryast/crud-14.git/database"
	"github.com/kryast/crud-14.git/models"
	"github.com/kryast/crud-14.git/router"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	db.AutoMigrate(&models.Product{})

	r := router.SetupRouter(db)
	r.Run(":8080")
}
