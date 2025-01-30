package hash

import (
	"crypto/rand"
	"fmt"
)

func GenerateRandomHash() string {
	var hashBytes = make([]byte, 16)
	if _, err := rand.Read(hashBytes); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", hashBytes)
}
