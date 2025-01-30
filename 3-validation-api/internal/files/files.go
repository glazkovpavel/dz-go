package files

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.filename)
	return file, err

}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close() // defer пушит выполнение ф-ции в конец стека вызова
	if err != nil {
		fmt.Println(err)
		return
	}
	color.Green("Запись произошла успешна")

}
