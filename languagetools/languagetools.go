package languagetools

import (
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func CalculateProbability(s string) map[rune]uint {
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

// TODO: add error return
func SwapCharactersInAlphabet(alphabet string, char1, char2 rune) string {
	newAlphabet := []rune(alphabet)
	index1 := FindIndexOfString(alphabet, char1)
	index2 := FindIndexOfString(alphabet, char2)
	newAlphabet[index1], newAlphabet[index2] = newAlphabet[index2], newAlphabet[index1]
	return string(newAlphabet)
}

func sortAlphabet(alphabetNormalOrdered, alphabetSecretOrdered, alphabetNormal string) (string, string) {
	secretAlphabet := ""
	for _, x := range alphabetNormal {
		index := FindIndexOfString(alphabetNormalOrdered, x)
		if index < len(alphabetSecretOrdered) {
			secretAlphabet += string(alphabetSecretOrdered[index])
		}
	}
	return alphabetNormal, secretAlphabet
}

func GetAlphabetsOrderProbability(textSecret, alphabetNormal string, alphabetNormalProb []Alphabet) (string, string) {
	frequenties := CalculateProbability(textSecret)
	encryptedAlphabet := []Alphabet{}
	for i, x := range frequenties {
		encryptedAlphabet = append(encryptedAlphabet, Alphabet{Character: string(i), Probability: float32(x)})
	}

	sort.Slice(alphabetNormalProb, func(i, j int) bool {
		return alphabetNormalProb[i].Probability > alphabetNormalProb[j].Probability
	})

	sort.Slice(encryptedAlphabet, func(i, j int) bool {
		return encryptedAlphabet[i].Probability > encryptedAlphabet[j].Probability
	})

	alphabetNormalOrdered := ""
	for _, y := range alphabetNormalProb {
		alphabetNormalOrdered += strings.ToLower(string(y.Character))
	}

	alphabetSecretOrdered := ""
	for _, y := range encryptedAlphabet {
		percentage := float64(y.Probability) / float64(len(textSecret)) * 100
		percentage = toFixed(percentage, 3)
		alphabetSecretOrdered += string(y.Character)
	}

	alphabetNormalOrdered, alphabetSecretOrdered = sortAlphabet(alphabetNormalOrdered, alphabetSecretOrdered, alphabetNormal)
	return alphabetNormalOrdered, alphabetSecretOrdered
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

func IndexOfCoincidence(str string) float64 {
	num := 0
	den := 0

	counst := CalculateProbability(str)

	for _, x := range counst {
		val := int(x)
		num += val * (val - 1)
		den += val
	}
	if den == 0 {
		return 0.0
	}
	return float64(num) / float64(float64(den)*float64(den-1))
}
