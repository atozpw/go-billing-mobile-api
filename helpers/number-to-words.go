package helpers

import (
	"fmt"
	"strings"
)

func NumberToWords(num int) string {
	var s string
	satuan := [12]string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan", "Sepuluh", "Sebelas"}
	if num < 12 {
		s = satuan[num]
	} else if num < 20 {
		s = fmt.Sprintf("%s Belas", NumberToWords(num-10))
	} else if num < 100 {
		s = fmt.Sprintf("%s Puluh %s", NumberToWords(num/10), NumberToWords(num%10))
	} else if num < 200 {
		s = fmt.Sprintf("Seratus %s", NumberToWords(num-100))
	} else if num < 1000 {
		s = fmt.Sprintf("%s Ratus %s", NumberToWords(num/100), NumberToWords(num%100))
	} else if num < 2000 {
		s = fmt.Sprintf("Seribu %s", NumberToWords(num-1000))
	} else if num < 1000000 {
		s = fmt.Sprintf("%s Ribu %s", NumberToWords(num/1000), NumberToWords(num%1000))
	} else if num < 2000000 {
		s = fmt.Sprintf("Satu Juta %s", NumberToWords(num-1000000))
	} else if num < 1000000000 {
		s = fmt.Sprintf("%s Juta %s", NumberToWords(num/1000000), NumberToWords(num%1000000))
	} else if num < 2000000000 {
		s = fmt.Sprintf("Satu Milyar %s", NumberToWords(num-1000000000))
	} else if num < 1000000000000 {
		s = fmt.Sprintf("%s Milyar %s", NumberToWords(num/1000000000), NumberToWords(num%1000000000))
	} else if num < 2000000000000 {
		s = fmt.Sprintf("Satu Triliun %s", NumberToWords(num-1000000000000))
	} else if num < 1000000000000000 {
		s = fmt.Sprintf("%s Triliun %s", NumberToWords(num/1000000000000), NumberToWords(num%1000000000000))
	}
	return strings.TrimSpace(s)
}
