package cesar

func shiftDecrypt(char rune, shift int) string {
	ascii := int(char)
	ascii -= shift

	if ascii >= ALPHABETSTART {
		return string(ascii)
	}

	ascii += ALPHABETEND
	return string(ascii)
}

func Decrypt(text string, shift int) string {
	newText := ""
	for _, char := range text {
		newCharacter := shiftDecrypt(char, shift)
		newText += newCharacter
	}
	return newText
}
