package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println(server.ListenAndServe())
}
