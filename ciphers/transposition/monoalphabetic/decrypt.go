package monoalphabetic

import (
	"strings"
)

func TransformDecrypt(char rune, alphabtNormal string, alphabetSecret string) string {
	index := strings.Index(alphabetSecret, string(char))
	decryptedCharacter := alphabtNormal[index]
	return string(decryptedCharacter)
}

func Decrypt(text string, alphabtNormal string, alphabetSecret string) string {
	if string(alphabtNormal[len(alphabtNormal)-1]) != string(" "){
		alphabtNormal += " "
	}
	if string(alphabetSecret[len(alphabetSecret)-1]) != string(" "){
		alphabetSecret += " "
	}

	newText := ""
	for _, char := range text {
		newText += TransformDecrypt(char, alphabtNormal, alphabetSecret)
	}
	return newText
}
