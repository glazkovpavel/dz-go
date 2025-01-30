package verify

import (
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	emailPkg "github.com/jordan-wright/email"
	"go/validation-api/configs"
	"go/validation-api/internal/email"
	"go/validation-api/pkg/request"
	"go/validation-api/pkg/response"
	"log"
	"net/http"
	"net/smtp"
)

type VerifierHandlerDeps struct {
	*configs.Config
	*email.EmailRepository
}

type VerifierHandler struct {
	*configs.Config
	*email.EmailRepository
}

func NewVerifierHandler(router *http.ServeMux, deps VerifierHandlerDeps) {
	handler := &VerifierHandler{
		Config:          deps.Config,
		EmailRepository: deps.EmailRepository,
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
		err = handler.AddEmail(*newEmail)
		if err != nil {
			color.Red(err.Error())
			return
		}
		err = handler.sendEmail(newEmail)
		if err != nil {
			color.Red(err.Error())
			return
		}

	}
}

func (handler *VerifierHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		var emails []email.Email
		isDeleted := false
		for _, emailData := range handler.Emails {
			if emailData.Hash != hash {
				emails = append(emails, emailData)
				continue
			}
			isDeleted = true
		}
		handler.Emails = emails
		err := handler.EmailRepository.Save()
		if err != nil {
			color.Red(err.Error())
		}
		data := VerifyResponse{
			VerificationPassed: isDeleted,
		}
		if isDeleted {
			response.Json(w, data, http.StatusOK)
			color.Green("Проверка пройдена")
			return
		}
		response.Json(w, data, http.StatusForbidden)
		color.Red("Проверка не пройдена")
	}
}

func (handler *VerifierHandler) sendEmail(emailData *email.Email) error {
	hash := emailData.Hash
	href := "http://localhost:8081/send/" + hash
	e := emailPkg.NewEmail()
	e.From = "glazprom@bk.ru"
	e.To = []string{emailData.Email}
	e.Subject = "Подтверждение регистрации"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(fmt.Sprintf("<h4>Подтвердите вашу почту, перейдя по ссылке</h4><br><a href=\"%s\">Нажмите сюда</a>", href))

	auth := smtp.PlainAuth(
		"",
		handler.Config.EmailConf.Address,
		handler.Config.EmailConf.Password,
		"smtp.mail.ru",
	)

	// Настройка TLS соединения
	tlsConfig := &tls.Config{
		ServerName: "smtp.mail.ru",
	}

	conn, err := tls.Dial("tcp", "smtp.mail.ru:465", tlsConfig)
	if err != nil {
		log.Fatalf("Ошибка при подключении к SMTP-серверу: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.mail.ru")
	if err != nil {
		log.Fatalf("Ошибка при создании SMTP-клиента: %v", err)
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		log.Fatalf("Ошибка при аутентификации: %v", err)
	}

	if err = client.Mail(e.From); err != nil {
		log.Fatalf("Ошибка при установке отправителя: %v", err)
	}

	for _, addr := range e.To {
		if err = client.Rcpt(addr); err != nil {
			log.Fatalf("Ошибка при добавлении получателя: %v", err)
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Fatalf("Ошибка при подготовке к передаче данных: %v", err)
	}
	defer w.Close()

	bytes, err := e.Bytes()
	if err != nil {
		log.Fatalf("Ошибка при получении байтов письма: %v", err)
	}

	_, err = w.Write(bytes)
	if err != nil {
		log.Fatalf("Ошибка при отправке письма: %v", err)
	}

	client.Quit()

	log.Println("Письмо успешно отправлено.")
	return nil
}
