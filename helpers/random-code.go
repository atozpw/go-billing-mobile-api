package helpers

import (
	"math/rand"
	"time"
)

func RandomCode(length int) string {

	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	letters := []rune("0123456789")

	code := make([]rune, length)

	for i := range code {
		code[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(code)

}
