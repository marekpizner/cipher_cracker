package monoalphabetic

import (
	"strings"
	"unicode"
)

func TransformDecrypt(char rune, alphabtNormal string, alphabetSecret string) string {
	index := strings.Index(alphabetSecret, string(char))
	decryptedCharacter := alphabtNormal[index]
	return string(decryptedCharacter)
}

func Decrypt(text string, alphabtNormal string, alphabetSecret string) string {
	newText := ""
	for _, char := range text {
		if unicode.IsSpace(char) {
			newText += string(char)
		} else {
			newText += TransformDecrypt(char, alphabtNormal, alphabetSecret)
		}
	}
	return newText
}
