package monoalphabetic

import (
	"strings"
	"unicode"
)

// !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}

func TransformEncrypt(char rune, alphabtNormal string, alphabetSecret string) string {
	index := strings.Index(alphabtNormal, string(char))
	encryptedCharacter := alphabetSecret[index]
	return string(encryptedCharacter)
}

func Encrypt(text string, alphabtNormal string, alphabetSecret string) string {
	newText := ""
	for _, char := range text {
		if unicode.IsSpace(char) {
			newText += string(char)
		} else {
			newText += TransformEncrypt(char, alphabtNormal, alphabetSecret)
		}
	}
	return newText
}
