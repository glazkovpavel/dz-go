package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/order-api/configs"
	"go/order-api/internal/product"
	"go/order-api/pkg/db"
	"go/order-api/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05", // Формат времени
	})

	productRepository := product.NewRepository(db)
	product.NewProductHandler(router, product.HandlerDeps{
		Repository: productRepository,
	})

	stack := middleware.Chain(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Listening on port 8081")
	server.ListenAndServe()
}
