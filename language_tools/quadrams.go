package language_tools

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseQuadrans(fileContent string) (map[string]int, int) {
	arrayFileContent := strings.Split(fileContent, "\n")
	arrayFileContent = arrayFileContent[:len(arrayFileContent)-1]

	quadgrams := make(map[string]int)
	maxValue := 0

	for _, x := range arrayFileContent {
		quadgram := strings.ToLower(strings.Split(x, " ")[0])
		value := strings.Split(x, " ")[1]
		intValue, _ := strconv.Atoi(value)
		quadgrams[quadgram] = intValue
		if intValue > maxValue {
			maxValue = intValue
		}
	}
	return quadgrams, maxValue
}

func CalculateQuadgrams(textSecret string, qLength int) map[string]float64 {
	cleantextSecret := strings.Replace(textSecret, " ", "", -1)
	quadgrams := make(map[string]float64)
	for i := 0; i < len(cleantextSecret)-qLength; i += 1 {
		quadgram := cleantextSecret[i : i+qLength]
		if val, ok := quadgrams[quadgram]; ok {
			quadgrams[quadgram]++
			val++
		} else {
			quadgrams[quadgram] = 1
		}
	}
	return quadgrams
}

func ReadQuadramsFromFile(path string) map[string]float64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	bString := string(b)
	parsedQuadrants, maxValue := parseQuadrans(bString)
	newParsedQuadgram := make(map[string]float64)

	for key, value := range parsedQuadrants {
		newValue := float64(value) / float64(maxValue) * 100
		newParsedQuadgram[strings.ToLower(key)] = newValue
	}
	return newParsedQuadgram
}
