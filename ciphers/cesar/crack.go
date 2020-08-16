package cesar

import (
	"fmt"
	"regexp/syntax"
)

func containEnglisWords(try_spaces []string) bool {
	total_perc := len(try_spaces)
	accuracy_perc := 0

	for _, word := range try_spaces {
		if syntax.IsWordChar([]rune(word)[0]) {
			accuracy_perc += 1
		}
	}
	percentage := float32(accuracy_perc) / float32(total_perc)
	if percentage > 0.8 {
		return true
	}
	return false

}

func Crack(text string) {

	for i := 1; i < 27; i++ {
		try := Decrypt(text, i)
		// try_spaces := strings.Split(try, " ")
		fmt.Println(try)

	}

}
