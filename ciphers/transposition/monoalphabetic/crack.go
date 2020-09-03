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

// TODO: refactor this

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func sortAlphabet(aNormal, aSecret, alphabetNormal string) (string, string) {
	secretAlphabet := ""
	for _, x := range alphabetNormal {
		index := language_tools.FindIndexOfString(aNormal, x)
		secretAlphabet += string(aSecret[index])
	}
	return alphabetNormal, secretAlphabet
}

func getAlphabetsOrderProbability(textSecret, alphabetNormal string, alphabetNormalProbability []language_tools.Alphabet) (string, string) {
	frequenties := check(textSecret)
	encryptedAlphabet := []language_tools.Alphabet{}
	for i, x := range frequenties {
		encryptedAlphabet = append(encryptedAlphabet, language_tools.Alphabet{Character: string(i), Probability: float32(x)})
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

func calculateLocalMaximum(textSecret, alphabetReal, alphabetSecret string, reaQuadgrams map[string]float64, repeatIterations int) (int, string) {
	maxFitnes := 0
	bestAlphabet := alphabetSecret
	alphabetSecretNew := alphabetSecret

	for c := 0; c < repeatIterations; c++ {
		for i := 0; i < len(alphabetSecret)-1; i++ {
			for j := i + 1; j < len(alphabetSecret); j++ {

				enc := Decrypt(textSecret, alphabetReal, alphabetSecretNew)
				encryptedQuadrams := language_tools.CalculateQuadgrams(enc)
				tmpFitnes := 0
				for keyE, _ := range encryptedQuadrams {
					if valueR, ok := reaQuadgrams[keyE]; ok {
						tmpFitnes += int(valueR)
					}
				}
				if tmpFitnes > maxFitnes {
					maxFitnes = tmpFitnes
					bestAlphabet = alphabetSecretNew
				}

				char1 := rune(bestAlphabet[i])
				char2 := rune(bestAlphabet[j])

				alphabetSecretNew = language_tools.SwapCharactersInAlphabet(bestAlphabet, char1, char2)
			}
		}
	}
	return maxFitnes, bestAlphabet
}

func Crack(textSecret, alphabetNormal string, realQuadgrams map[string]float64, alphabetNormalProbability []language_tools.Alphabet) string {
	textSecret = strings.ToLower(textSecret)
	alphabetNormal = strings.ToLower(alphabetNormal)
	alphabetNormal = strings.ReplaceAll(alphabetNormal, " ", "")
	alphabetNormal, alphabetSecret := getAlphabetsOrderProbability(textSecret, alphabetNormal, alphabetNormalProbability)

	bestScore := 0
	bestScoreHits := 0
	consolidate := 2
	repeatIterations := 1

	bestAlphabet := alphabetSecret

	for i := 0; i < 1000; i++ {

		score, alphabetSecretNew := calculateLocalMaximum(textSecret, alphabetNormal, bestAlphabet, realQuadgrams, repeatIterations)
		if score > bestScore {
			bestScore = score
			bestAlphabet = alphabetSecretNew
		} else if score == bestScore && score != 0 {
			bestScoreHits++
			if bestScoreHits == consolidate {
				break
			}
			continue
		}
	}

	enc := Decrypt(textSecret, alphabetNormal, bestAlphabet)
	fmt.Println(strings.ToLower(alphabetNormal))
	fmt.Println(bestAlphabet)
	fmt.Println(enc)
	return enc

}
