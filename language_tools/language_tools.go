package language_tools

import (
	"math/rand"
	"strings"
	"time"
)

func Shuffle(src string) string {
	src = strings.ReplaceAll(src, " ", "")
	final := []byte(src)
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return string(final) + " "
}

func FindIndexOfString(str string, char rune) int {
	for i, x := range str {
		if string(x) == string(char) {
			return i
		}
	}
	return -1
}

func SwapCharactersInAlphabet(alphabet string, char1, char2 rune) string {
	newAlphabet := []rune(alphabet)
	index1 := FindIndexOfString(alphabet, char1)
	index2 := FindIndexOfString(alphabet, char2)
	newAlphabet[index1], newAlphabet[index2] = newAlphabet[index2], newAlphabet[index1]
	return string(newAlphabet)
}
