package main

import (
	"fmt"
	"go/validation-api/configs"
	"go/validation-api/internal/email"
	"go/validation-api/internal/files"
	"go/validation-api/internal/verify"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	emailsWithDb := email.NewEmailWithDb(files.NewJsonDb("db/data.json"))
	verify.NewVerifierHandler(router, verify.VerifierHandlerDeps{Config: conf, EmailsWithDb: emailsWithDb})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Listening on port 8081")
	server.ListenAndServe()
}
