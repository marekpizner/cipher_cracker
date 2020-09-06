package vigenere

import (
	"fmt"
	"sort"

	"github.com/khan745/cipher_cracker/language_tools"
)

type kv struct {
	Key   string
	Value float64
}

type ka struct {
	Key   string
	Value []int
}

func orderArr(data map[string][]int) []ka {
	var ss []ka
	for k, v := range data {
		ss = append(ss, ka{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return len(ss[i].Value) > len(ss[j].Value)
	})

	return ss
}

func orderdic(data map[string]float64) []kv {
	var ss []kv
	for k, v := range data {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func Crack(textSecret string, realQuadgrams map[string]float64, alphabetNormalProbability []language_tools.Alphabet, alphabets []string) {
	//TODO: crack viniger cipher
	// 1. find key length
	// 2. frequency analysis
	keyLenths := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, k := range keyLenths {
		history := language_tools.RepetativeStrings(textSecret, k)
		historyOrder := orderArr(history)
		fmt.Println(k)
		for _, x := range historyOrder[:10] {
			fmt.Println(x)
		}
	}

	// fmt.Println(quadgramsSecrets)

}
