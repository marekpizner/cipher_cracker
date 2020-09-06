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

func spacing(array []int) []int {
	var spacings []int

	for i := 1; i < len(array); i++ {
		diff := array[i] - array[i-1]
		spacings = append(spacings, diff)
		if len(spacings) > 1 {
			spacings = append(spacings, spacings[len(spacings)-2]+diff)
		}
	}
	sort.Ints(spacings[:])
	return spacings
}

func Crack(textSecret string, realQuadgrams map[string]float64, alphabetNormalProbability []language_tools.Alphabet, alphabets []string) {
	//TODO: crack viniger cipher
	// 1. find key length
	// 2. frequency analysis
	// textSecret = "PPQCAXQVEKGYBNKMAZUYBNGBALJONITSZMJYIMVRAGVOHTVRAUCTKSGDDWUOXITLAZUVAVVRAZCVKBQPIWPOU"

	keyLenths := []int{3, 4, 5, 6}

	for _, k := range keyLenths {
		history := language_tools.RepetativeStrings(textSecret, k)
		historyOrder := orderArr(history)
		fmt.Println(k)
		for _, x := range historyOrder[:] {
			if len(x.Value) > 2 {
				x.Value = spacing(x.Value)
				fmt.Println(x)
			}
		}
	}

}
