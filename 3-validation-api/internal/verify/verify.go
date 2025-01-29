package verify

import (
	"fmt"
	"go/validation-api/configs"
	"go/validation-api/internal/email"
	"go/validation-api/pkg/request"
	"net/http"
)

type VerifierHandlerDeps struct {
	*configs.Config
	*email.EmailsWithDb
}

type VerifierHandler struct {
	*configs.Config
	*email.EmailsWithDb
}

func NewVerifierHandler(router *http.ServeMux, deps VerifierHandlerDeps) {
	handler := &VerifierHandler{
		Config:       deps.Config,
		EmailsWithDb: deps.EmailsWithDb,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/send/{hash}", handler.Verify())
}

func (handler *VerifierHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[EmailRequest](&w, req)
		if err != nil {
			return
		}

		fmt.Println(*body)

		newEmail := email.NewEmail(body.Email)
		handler.AddEmail(*newEmail)

	}
}

func (handler *VerifierHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
