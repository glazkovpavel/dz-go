package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	newHandler(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()

}
