package main

import (
	"fmt"
	"go/validation-api/configs"
	"go/validation-api/internal/verify"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifierHandler(router, verify.VerifierHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Listening on port 8081")
	server.ListenAndServe()
}
