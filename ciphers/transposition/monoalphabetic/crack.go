package monoalphabetic

import (
	"fmt"
	"math"
	"sort"

	"github.com/khan745/cipher_cracker/language_tools"
)

func check(s string) map[rune]uint {
	m := make(map[rune]uint, len(s))
	for _, r := range s {
		if r != ' ' {
			m[r]++
		}
	}
	return m
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func Crack(text string) {
	frequenties := check(text)
	encryptedAlphabet := []language_tools.Alphabet{}
	for i, x := range frequenties {
		encryptedAlphabet = append(encryptedAlphabet, language_tools.Alphabet{Character: string(i), Probability: float32(x)})
	}

	alphabetRealProb := language_tools.ReadFiles("./alphabets", "csv")

	sort.Slice(alphabetRealProb, func(i, j int) bool {
		return alphabetRealProb[i].Probability > alphabetRealProb[j].Probability
	})

	sort.Slice(encryptedAlphabet, func(i, j int) bool {
		return encryptedAlphabet[i].Probability > encryptedAlphabet[j].Probability
	})

	alphabetReal := ""
	for x, y := range alphabetRealProb {
		fmt.Println(x, y)

		alphabetReal += string(y.Character)

	}

	alphabetSecret := ""
	for x, y := range encryptedAlphabet {
		percentage := float64(y.Probability) / float64(len(text)) * 100
		percentage = toFixed(percentage, 3)
		fmt.Println(x, y.Character, percentage)
		if y.Character != string(' ') {
			alphabetSecret += string(y.Character)
		}
	}

	fmt.Println(alphabetReal)
	fmt.Println(alphabetSecret)

	enc := Decrypt(text, alphabetReal, alphabetSecret)
	fmt.Println(enc)

}
