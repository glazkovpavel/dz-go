package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/internal/product"
	"go/order-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	productRepository := product.NewRepository(db)
	product.NewProductHandler(router, product.HandlerDeps{
		Repository: productRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Listening on port 8081")
	server.ListenAndServe()
}
