package iteration

import (
	"strings"
)

func Reverse(word string) string {
	chars := strings.Split(word, "")
	var reversed string

	for i := len(chars) - 1; i >= 0; i-- {
		reversed += chars[i]
	}

	return reversed
}
