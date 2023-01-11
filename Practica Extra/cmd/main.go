package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/jmarambio/prueba/cmd/routes"
	"github.com/jmarambio/prueba/pkg/store"
)

func main() {
	_ = godotenv.Load()
	storage := store.NewStore("./db/db.json")
	en := gin.Default()
	r := routes.NewRouter(en, storage)
	r.SetRoutes()
	if err := en.Run(); err != nil {
		log.Fatal(err)
	}

	/*
		router := gin.Default()
		products := router.Group("/products")
		{
			products.GET("", handler.GetProducts)
			products.GET(":id", handler.GetProductById)
			products.GET("/search", handler.GetProductByFilter)
			products.POST("", handler.AddProduct)
		}

		router.Run()*/
}
