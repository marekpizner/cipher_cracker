package cesar

func shiftDecrypt(char rune, shift int) rune {
	ascii := int(char) - shift

	if ascii < ALPHABETSTART {
		return rune(ALPHABETEND - (ALPHABETSTART - ascii) + 1)
	}

	return rune(ascii)
}

func Decrypt(text string, shift int) string {
	newText := ""
	for _, char := range text {
		newCharacter := shiftDecrypt(char, shift)
		newText += string(newCharacter)
	}
	return newText
}
