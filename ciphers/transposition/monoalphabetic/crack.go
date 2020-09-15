package monoalphabetic

import (
	"strings"

	"github.com/khan745/cipher_cracker/languagetools"
)

func calculateLocalMaximum(textSecret, alphabetReal, alphabetSecret string, reaQuadgrams map[string]float64, repeatIterations, qLength int) (int, string) {
	maxFitnes := 0
	bestAlphabet := alphabetSecret
	alphabetSecretNew := alphabetSecret

	for c := 0; c < repeatIterations; c++ {
		for i := 0; i < len(alphabetSecret)-1; i++ {
			for j := i + 1; j < len(alphabetSecret); j++ {

				enc := Decrypt(textSecret, alphabetReal, alphabetSecretNew)
				encryptedQuadrams := languagetools.CalculateQuadgrams(enc, qLength)
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

				alphabetSecretNew = languagetools.SwapCharactersInAlphabet(bestAlphabet, char1, char2)
			}
		}
	}
	return maxFitnes, bestAlphabet
}

func Crack(textSecret, alphabetNormal string, realQuadgrams map[string]float64, alphabetNormalProbability []languagetools.Alphabet) (string, string) {
	textSecret = strings.ToLower(textSecret)
	alphabetNormal = strings.ToLower(alphabetNormal)
	alphabetNormal = strings.ReplaceAll(alphabetNormal, " ", "")
	alphabetNormal, alphabetSecret := languagetools.GetAlphabetsOrderProbability(textSecret, alphabetNormal, alphabetNormalProbability)

	bestScore := 0
	bestScoreHits := 0
	consolidate := 2
	repeatIterations := 1
	qLength := 4

	bestAlphabet := alphabetSecret

	for i := 0; i < 1000; i++ {

		score, alphabetSecretNew := calculateLocalMaximum(textSecret, alphabetNormal, bestAlphabet, realQuadgrams, repeatIterations, qLength)
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
	return enc, bestAlphabet

}
