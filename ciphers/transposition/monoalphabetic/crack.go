package monoalphabetic

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"math/rand"
	"time"
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
		alphabetReal += string(y.Character)
	}

	alphabetSecret := ""
	for _, y := range encryptedAlphabet {
		percentage := float64(y.Probability) / float64(len(text)) * 100
		percentage = toFixed(percentage, 3)
		alphabetSecret += string(y.Character)
	}
	alphabetReal += " "
	alphabetSecret += " "
	return alphabetReal, alphabetSecret
}

func CalculateQuadgrams(text string) map[string]float64 {
	cleanText := strings.Replace(text, " ", "", -1)
	quadgrams := make(map[string]float64)
	for i := 0; i < len(cleanText)-4; i+=1 {
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
	Value float64
}


func orderdic(data map[string]float64) []kv{
	var ss []kv
	for k, v := range data {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}



func shuffle(src string) string {
	src = strings.ReplaceAll(src, " ", "")
	final := []byte(src)
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return string(final) + " "
}

func findIndexOfString(str string, char rune) int{
	for i,x := range str {
		if string(x) == string(char){
			return i
		}
	}
	return -1
}

func swapCharactersInAlphabet(alphabet string, char1, char2 rune) string{
	newAlphabet := []rune(alphabet)
	index1 := findIndexOfString(alphabet, char1)
	index2 := findIndexOfString(alphabet, char2)
	newAlphabet[index1], newAlphabet[index2] = newAlphabet[index2], newAlphabet[index1]
	// fmt.Println(alphabet, string(char1), string(char2), string(newAlphabet))
	return string(newAlphabet)
}

func calculateLocalMaximum(text, alphabetReal, alphabetSecret string, reaQuadgrams map[string]float64) (int, string){
	maxFitnes := 0
	bestAlphabet := alphabetSecret
	alphabetSecretNew := alphabetSecret

	for i:=0; i<len(alphabetReal)-1;i++{
		for j:=i+1; j<len(alphabetReal);j++{

			char1 := rune(bestAlphabet[i])
			char2 := rune(bestAlphabet[j])
			
			alphabetSecretNew = swapCharactersInAlphabet(bestAlphabet, char1, char2)


			enc := Decrypt(text, alphabetReal, alphabetSecretNew)
			encryptedQuadrams := CalculateQuadgrams(enc)
			tmpFitnes := 0
			for key_e, _ := range encryptedQuadrams{
				if value_r, ok := reaQuadgrams[key_e]; ok {
					tmpFitnes += int(value_r)
				}
			}

			if tmpFitnes > maxFitnes {
				maxFitnes = tmpFitnes
				bestAlphabet = alphabetSecretNew
				// fmt.Println(bestAlphabet, maxFitnes)
			}else{
				alphabetSecretNew = swapCharactersInAlphabet(bestAlphabet, char1, char2)
			}
		}
	}
	return maxFitnes, bestAlphabet
}

func Crack(text string) {
	reaQuadgrams := language_tools.ReadQuadrams("./english_quadgrams.txt")
	alphabetReal, alphabetSecret := getAlphabets(text)

	alphabetSecret = strings.ReplaceAll(alphabetSecret, " ", "")
	alphabetReal = strings.ReplaceAll(alphabetReal, " ", "")

	bestScore := 0
	bestScoreHits := 0
	consolidate := 2

	bestAlphabet := alphabetSecret
	alphabetSecretNew := alphabetSecret

	fmt.Println(alphabetSecret)
	
	for i:=0; i<1000;i++{
			alphabetSecretNew = shuffle(alphabetSecret)
			// fmt.Println(alphabetSecretNew)
			score, alphabetSecretNew := calculateLocalMaximum(text,alphabetReal,alphabetSecretNew,reaQuadgrams)

			if score > bestScore {
				// fmt.Print("\033[G\033[K")
				bestScore = score
				bestAlphabet = alphabetSecretNew
				fmt.Println(bestAlphabet, bestScore)
				// fmt.Print("\033[A") 
			}else if score == bestScore {
				bestScoreHits ++
				if bestScoreHits == consolidate{
					break
				}
			}
	}

	enc := Decrypt(text, alphabetReal, bestAlphabet)
	fmt.Println(strings.ToLower(alphabetReal))
	fmt.Println(bestAlphabet)
	fmt.Println(enc)

}
