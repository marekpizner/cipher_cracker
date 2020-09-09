package vigenere

import (
	"unicode"

	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"
)

func transformEncrypt(char rune, alphabets []string, keyWorldCharacter rune) string {

	alphabetConcrete := alphabets[0]
	for _, i := range alphabets {
		if rune(i[0]) == keyWorldCharacter {
			alphabetConcrete = i
		}
	}
	// fmt.Println("Encrypting: ", string(char), " with alphabet: ", alphabetConcrete, " with keyword: ", string(keyWorldCharacter))
	encryptedCharacter := monoalphabetic.TransformEncrypt(char, alphabets[0], alphabetConcrete)
	return encryptedCharacter
}

func Encrypt(text string, alphabet []string, keyWord string) string {
	newText := ""
	ik := 0
	for _, char := range text {
		if unicode.IsSpace(char) {
			newText += string(char)
			continue
		}
		keywordCharacterIndex := ik % len(keyWord)
		keywordCharacter := rune(keyWord[keywordCharacterIndex])
		newText += transformEncrypt(char, alphabet, keywordCharacter)
		ik++
	}
	return newText
}
