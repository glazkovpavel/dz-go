package email

import (
	"encoding/json"
	"github.com/fatih/color"
	"go/validation-api/pkg/hash"
)

type EmailSt struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type Vault struct {
	Emails []EmailSt `json:"emails"`
}

type EmailsWithDb struct {
	Vault
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

func NewEmail(email string) *EmailSt {
	return &EmailSt{
		Email: email,
		Hash:  hash.GenerateRandomHash(),
	}
}

func NewEmailWithDb(db Db) *EmailsWithDb {
	file, err := db.Read()
	if err != nil {
		return &EmailsWithDb{
			Vault: Vault{
				Emails: []EmailSt{},
			},
			db: db,
		}
	}
	var emails Vault
	err = json.Unmarshal(file, &emails)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json", err.Error())
		return &EmailsWithDb{
			Vault: Vault{
				Emails: []EmailSt{},
			},
			db: db,
		}
	}
	return &EmailsWithDb{
		Vault: emails,
		db:    db,
	}
}

func (emails *EmailsWithDb) AddEmail(email EmailSt) error {
	emails.Emails = append(emails.Emails, email)
	err := emails.save()
	if err != nil {
		return err
	}
	return nil

}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	return file, err
}

func (emails *EmailsWithDb) save() error {
	data, err := emails.Vault.ToBytes()
	//encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Не удалось преобразовать", err.Error())
		return err
	}
	emails.db.Write(data)
	return nil
}
