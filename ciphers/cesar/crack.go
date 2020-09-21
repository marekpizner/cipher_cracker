package cesar

import (
	"fmt"
	"regexp/syntax"
)

func containEnglisWords(trySpaces []string) bool {
	totalPerc := len(trySpaces)
	accuracyPerc := 0

	for _, word := range trySpaces {
		if syntax.IsWordChar([]rune(word)[0]) {
			accuracyPerc++
		}
	}
	percentage := float32(accuracyPerc) / float32(totalPerc)
	if percentage > 0.8 {
		return true
	}
	return false

}

func Crack(text string) {

	for i := 1; i < 27; i++ {
		try := Decrypt(text, i)
		fmt.Println(try)
	}

}
