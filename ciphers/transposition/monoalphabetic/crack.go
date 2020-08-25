package monoalphabetic

import (
	"fmt"
	"math"
	"sort"
	"strings"

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

func getAlphabets(text string) (string, string) {
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
	for _, y := range alphabetRealProb {
		fmt.Println(y.Character, y.Probability)
		alphabetReal += string(y.Character)
	}

	alphabetSecret := ""
	for _, y := range encryptedAlphabet {
		percentage := float64(y.Probability) / float64(len(text)) * 100
		percentage = toFixed(percentage, 3)
		fmt.Println(y.Character, percentage)
		alphabetSecret += string(y.Character)
	}
	alphabetReal += " "
	alphabetSecret += " "
	return alphabetReal, alphabetSecret
}

func CalculateQuadgrams(text string) map[string]int {
	cleanText := strings.Replace(text, " ", "", -1)
	quadgrams := make(map[string]int)
	for i := 0; i < len(cleanText)-4; i++ {
		quadgram := cleanText[i : i+4]
		if val, ok := quadgrams[quadgram]; ok {
			quadgrams[quadgram]++
			val++
		} else {
			quadgrams[quadgram] = 1
		}
	}
	return quadgrams
}

type kv struct {
	Key   string
	Value int
}

func orderdic(data map[string]int) []kv{
	var ss []kv
	for k, v := range data {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func Crack(text string) {
	alphabetReal, alphabetSecret := getAlphabets(text)

	fmt.Println("-----------GuestAlphabet------------")
	fmt.Println(alphabetReal)
	fmt.Println(alphabetSecret)

	// enc := Decrypt(text, alphabetReal, alphabetSecret)
	// fmt.Println(enc)

	reaQuadgrams := language_tools.ReadQuadrams("./english_quadgrams.txt")
	secQuadgrams := CalculateQuadgrams(text)

	orderedRealQuadrams := orderdic(reaQuadgrams)
	orderedSecQuadgrams := orderdic(secQuadgrams)

	fmt.Println("--------------Real-------------------")
	for _,x := range orderedRealQuadrams[:10]{
		fmt.Println(x)
	}
	fmt.Println("--------------Sec-------------------")
	for _,x := range orderedSecQuadgrams[:10]{
		fmt.Println(x)
	}

	
}
