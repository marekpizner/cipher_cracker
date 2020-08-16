package vigenere

import (
	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"
)

func transformDecrypt(char rune, alphabet []string, keyWorldCharacter rune) string {

	alphabetConcrete := alphabet[0]
	for _, i := range alphabet {
		if rune(i[0]) == keyWorldCharacter {
			alphabetConcrete = i
		}
	}

	decryptedCharacter := monoalphabetic.TransformDecrypt(char, alphabet[0], alphabetConcrete)
	return decryptedCharacter
}

func Decrypt(text string, alphabet []string, keyWord string) string {
	newText := ""
	for i, char := range text {
		keywordCharacterIndex := i % len(keyWord)
		keywordCharacter := rune(keyWord[keywordCharacterIndex])
		newText += transformDecrypt(char, alphabet, keywordCharacter)
	}
	return newText
}
