package cesar

import "fmt"

const alphabet = []string{"a", "b"}

func Encrypt(text string, shift int) string {
	new_text := ""

	for _, char := range text {
		fmt.Print(char)
		char += 3
		fmt.Print("->")
		fmt.Print(char)
		new_text = new_text + string(char)
		fmt.Println()
	}
	return new_text
}
