package language_tools

import (
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func calculateProbability(s string) map[rune]uint {
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

func sortAlphabet(aNormal, aSecret, alphabetNormal string) (string, string) {
	secretAlphabet := ""
	for _, x := range alphabetNormal {
		index := FindIndexOfString(aNormal, x)
		secretAlphabet += string(aSecret[index])
	}
	return alphabetNormal, secretAlphabet
}

func GetAlphabetsOrderProbability(textSecret, alphabetNormal string, alphabetNormalProbability []Alphabet) (string, string) {
	frequenties := calculateProbability(textSecret)
	encryptedAlphabet := []Alphabet{}
	for i, x := range frequenties {
		encryptedAlphabet = append(encryptedAlphabet, Alphabet{Character: string(i), Probability: float32(x)})
	}

	sort.Slice(alphabetNormalProbability, func(i, j int) bool {
		return alphabetNormalProbability[i].Probability > alphabetNormalProbability[j].Probability
	})

	sort.Slice(encryptedAlphabet, func(i, j int) bool {
		return encryptedAlphabet[i].Probability > encryptedAlphabet[j].Probability
	})

	alphabetReal := ""
	for _, y := range alphabetNormalProbability {
		alphabetReal += strings.ToLower(string(y.Character))
	}

	alphabetSecret := ""
	for _, y := range encryptedAlphabet {
		percentage := float64(y.Probability) / float64(len(textSecret)) * 100
		percentage = toFixed(percentage, 3)
		alphabetSecret += string(y.Character)
	}
	alphabetReal, alphabetSecret = sortAlphabet(alphabetReal, alphabetSecret, alphabetNormal)
	return alphabetReal, alphabetSecret
}

func RepetativeStrings(text string, length int) map[string][]int {
	text = strings.ReplaceAll(text, " ", "")
	history := make(map[string][]int)

	for i := 0; i < len(text)-length; i++ {
		str := text[i : i+length]
		// fmt.Println(str)
		// fmt.Println(string(x))
		history[str] = append(history[str], i)
	}

	return history
}
