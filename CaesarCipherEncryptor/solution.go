package CaesarCipherEncryptor

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func CaesarCipherEncryptor(str string, key int) string {
	var builder strings.Builder

	alphabetLen := len(alphabet) - 1

	for _, s := range str {
		idx := strings.IndexRune(alphabet, s)
		if idx+key > alphabetLen {
			idx = ((idx - alphabetLen) + key) - 1
			builder.WriteByte(alphabet[idx])
		} else {
			builder.WriteByte(alphabet[idx+key])
		}
	}

	return builder.String()
}
