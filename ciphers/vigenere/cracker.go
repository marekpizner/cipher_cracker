package vigenere

import (
	"fmt"

	"github.com/khan745/cipher_cracker/language_tools"
)

func Crack(textSecret string, realQuadgrams map[string]float64, alphabetNormalProbability []language_tools.Alphabet, alphabets []string) {
	//TODO: crack viniger cipher
	// 1. find key length
	// 2. frequency analysis

	quadgramsSecrets := language_tools.CalculateQuadgrams(textSecret, 5)
	fmt.Println(quadgramsSecrets)

}
