package verify

import (
	"go/validation-api/configs"
	"net/http"
)

type VerifierHandlerDeps struct {
	*configs.Config
}

type VerifierHandler struct {
	*configs.Config
}

func NewVerifierHandler(router *http.ServeMux, deps VerifierHandlerDeps) {
	handler := &VerifierHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/send/{hash}", handler.Verify())
}

func (handler *VerifierHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (handler *VerifierHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
