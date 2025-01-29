package helpers

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CurrencyFormat(num int) string {
	em := message.NewPrinter(language.English)
	var enNumber string = em.Sprintf("%d", num)
	return enNumber
}
