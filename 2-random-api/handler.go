package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Handler struct{}

func newHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("/write", handler.getBytes())
}

func (handler *Handler) getBytes() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		randomInt := rune(rand.Intn(6) + 1)
		fmt.Println(randomInt)
		_, err := writer.Write([]byte(string(randomInt)))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
