package cesar

const ALPHABETSTART = 32
const ALPHABETEND = 126

func shiftEncrypt(char rune, shift int) string {
	ascii := int(char)
	ascii += shift

	if ascii <= ALPHABETEND {
		return string(ascii)
	}

	ascii -= ALPHABETEND
	return string(ascii)
}

func Encrypt(text string, shift int) string {
	newText := ""
	for _, char := range text {
		newCharacter := shiftEncrypt(char, shift)
		newText += newCharacter
	}
	return newText
}
