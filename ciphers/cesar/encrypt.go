package cesar

import "unicode"

const ALPHABETSTART = 97
const ALPHABETEND = 122

func shiftEncrypt(char rune, shift int) rune {
	ascii := int(char) + shift

	if ascii > ALPHABETEND {
		return rune(ALPHABETSTART + (ascii - ALPHABETEND) - 1)
	}

	return rune(ascii)
}

func Encrypt(text string, shift int) string {
	newText := ""
	for _, char := range text {
		if unicode.IsSpace(char) {
			newText += " "
			continue
		}
		newCharacter := shiftEncrypt(char, shift)
		newText += string(newCharacter)
	}
	return newText
}
