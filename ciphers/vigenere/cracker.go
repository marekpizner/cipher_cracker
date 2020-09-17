package vigenere

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"

	"github.com/khan745/cipher_cracker/languagetools"
)

type kv struct {
	Key   string
	Value float64
}

type kvv struct {
	Key   rune
	Value uint
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

func orderdicUint(data map[rune]uint) []kvv {
	var ss []kvv
	for k, v := range data {
		ss = append(ss, kvv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
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
func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	sort.Ints(result[:])
	return result
}

func calculateFactors(number, maxKeyLength int) []int {
	var factors []int

	if number < 2 {
		return factors
	}

	for i := 2; i < maxKeyLength+1; i++ {
		if number%i == 0 {
			factors = append(factors, i)
			otherFactor := number / i
			if otherFactor < maxKeyLength+1 && otherFactor != 1 {
				factors = append(factors, otherFactor)
			}
		}
	}

	factors = removeDuplicates(factors)
	return factors
}

func max(a []int) int {
	max := a[0]

	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

func countOccurence(array []int) map[int]int {
	dic := make(map[int]int)
	for _, x := range array {
		dic[x]++
	}
	return dic
}

func getNstring(str string, start, n int) string {
	newString := ""
	str = strings.ReplaceAll(str, " ", "")

	for i := 0; i < len(str)-start; i++ {

		if (i)%n == 0 {
			newString += string(str[start+i])
		}
	}
	return newString
}

func findLocalMaximum(bestKey, textSecret string, realQuadgrams map[string]float64, alphabets []string) string {
	qLength := 4
	maxFitnes := 0
	key := bestKey
	for c := 0; c < 2; c++ {
		for keyl := 0; keyl < len(bestKey); keyl++ {
			for al := 0; al < len(alphabets); al++ {

				dec := Decrypt(textSecret, alphabets, key)
				encryptedQuadrams := languagetools.CalculateQuadgrams(dec, qLength)

				tmpFitnes := 0
				for keyE, _ := range encryptedQuadrams {
					if valueR, ok := realQuadgrams[keyE]; ok {
						tmpFitnes += int(valueR)
					}
				}
				if tmpFitnes > maxFitnes {
					maxFitnes = tmpFitnes
					bestKey = key
				}

				key = bestKey[:keyl] + string(alphabets[al][0]) + bestKey[keyl+1:]
			}
		}
	}
	return bestKey
}

func Crack(textSecret string, realQuadgrams map[string]float64, alphabetNormalProb []languagetools.Alphabet, alphabets []string) {
	//TODO: crack viniger cipher
	// 1. find key length
	// 2. frequency analysis
	// textSecret = "PPQCAXQVEKGYBNKMAZUYBNGBALJONITSZMJYIMVRAGVOHTVRAUCTKSGDDWUOXITLAZUVAVVRAZCVKBQPIWPOU"
	englishIc := 0.0667
	keyLenths := []int{3, 4, 5, 6, 7, 8, 9, 10}
	maxKeyLength := max(keyLenths)
	posibleKeyLengths := make(map[int]int)

	for _, k := range keyLenths {
		history := languagetools.RepetativeStrings(textSecret, k)
		historyOrder := orderArr(history)
		allValues := []int{}
		fmt.Println("key length: ", k)
		for _, x := range historyOrder[:] {
			if len(x.Value) > 2 {
				x.Value = spacing(x.Value)
				allValues = append(allValues, x.Value...)
			}
		}
		allValues = removeDuplicates(allValues)
		fmt.Println("Spacing: ", allValues)

		allfactors := []int{}
		for _, x := range allValues {
			factors := calculateFactors(x, maxKeyLength)
			allfactors = append(allfactors, factors...)
		}
		// allfactors = removeDuplicates(allfactors)
		// fmt.Println("Factors: ", allfactors)
		occurences := countOccurence(allfactors)
		// fmt.Println("Occur: ", occurences)

		for k, v := range occurences {
			posibleKeyLengths[k] += v
		}
	}

	fmt.Println("Possible key lengths: ", posibleKeyLengths)
	// TODO
	// for each key length calculate key
	t := getNstring(textSecret, 0, 2)
	theBestGuseKeyLength := 2
	theBestGursKeyValue := math.Abs(englishIc - float64(languagetools.IndexOfCoincidence(t)))

	for k, _ := range posibleKeyLengths {
		t := getNstring(textSecret, 0, k)
		ic := languagetools.IndexOfCoincidence(t)
		fmt.Println("Key length: ", k, "IC: ", ic)
		if (englishIc - float64(ic)) < theBestGursKeyValue {
			theBestGuseKeyLength = k
			theBestGursKeyValue = math.Abs(englishIc - float64(ic))
		}
	}
	fmt.Println("Best key length: ", theBestGuseKeyLength, " with score: ", theBestGursKeyValue)
	key := ""
	for i := 0; i < theBestGuseKeyLength; i++ {
		t := getNstring(textSecret, i, theBestGuseKeyLength)
		fmt.Println(i, t)

		for _, alphabet := range alphabets {
			dec := monoalphabetic.Decrypt(t, alphabets[0], alphabet)

			prob := languagetools.CalculateProbability(dec)
			probOrder := orderdicUint(prob)
			if probOrder[0].Key == 'e' {
				fmt.Println(string(probOrder[0].Key), probOrder[0].Value, string(alphabet[0]))
				key += string(alphabet[0])
			}
		}
	}
	fmt.Println("Best key length: ", theBestGuseKeyLength, " with score: ", theBestGursKeyValue)
	fmt.Println("Best key:", key)
	bestKey := findLocalMaximum(key, textSecret, realQuadgrams, alphabets)
	fmt.Println("Best key:", bestKey)

}
