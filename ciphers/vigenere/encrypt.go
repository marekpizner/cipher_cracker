package vigenere

import (
	"fmt"

	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"
)

func transformEncrypt(char rune, alphabet []string, keyWorldCharacter rune) string {

	alphabetConcrete := alphabet[0]
	for _, i := range alphabet {
		if rune(i[0]) == keyWorldCharacter {
			alphabetConcrete = i
		}
	}
	fmt.Println(string(char), string(keyWorldCharacter), alphabetConcrete)
	encryptedCharacter := monoalphabetic.TransformEncrypt(char, alphabet[0], alphabetConcrete)
	return encryptedCharacter
}

func Encrypt(text string, alphabet []string, keyWord string) string {
	newText := ""
	for i, char := range text {
		keywordCharacterIndex := i % len(keyWord)
		keywordCharacter := rune(keyWord[keywordCharacterIndex])
		newText += transformEncrypt(char, alphabet, keywordCharacter)
	}
	return newText
}
