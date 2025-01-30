package email

import (
	"encoding/json"
	"github.com/fatih/color"
	"go/validation-api/pkg/hash"
)

type Email struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type EmailStore struct {
	Emails []Email `json:"emails"`
}

type EmailRepository struct {
	EmailStore
	db Db
}

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteWriter
	ByteReader
}

func NewEmail(email string) *Email {
	return &Email{
		Email: email,
		Hash:  hash.GenerateRandomHash(),
	}
}

func NewEmailWithDb(db Db) *EmailRepository {
	file, err := db.Read()
	if err != nil {
		return &EmailRepository{
			EmailStore: EmailStore{
				Emails: []Email{},
			},
			db: db,
		}
	}
	var emails EmailStore
	err = json.Unmarshal(file, &emails)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json", err.Error())
		return &EmailRepository{
			EmailStore: EmailStore{
				Emails: []Email{},
			},
			db: db,
		}
	}
	return &EmailRepository{
		EmailStore: emails,
		db:         db,
	}
}

func (emails *EmailRepository) AddEmail(email Email) error {
	emails.Emails = append(emails.Emails, email)
	err := emails.Save()
	if err != nil {
		return err
	}
	return nil

}

func (vault *EmailStore) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	return file, err
}

func (emails *EmailRepository) Save() error {
	data, err := emails.EmailStore.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать", err.Error())
		return err
	}
	emails.db.Write(data)
	return nil
}
