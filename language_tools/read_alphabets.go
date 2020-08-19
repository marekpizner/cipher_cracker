package language_tools

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func getFiles(path string, extension string) ([]string, error) {
	files, err := walkMatch(path, extension)
	if err != nil {
		return nil, err
	} else {
		return files, nil
	}
}

func readFile(path string) []Alphabet {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	alphabet := []Alphabet{}

	for scanner.Scan() {
		row := scanner.Text()
		char := strings.Split(row, ",")[0]
		probability_value, err := strconv.ParseFloat(strings.Split(row, ",")[1], 32)
		if err != nil {
			log.Fatal(err)
		}
		probability := float32(probability_value)
		alphabet = append(alphabet, Alphabet{Character: char, Probability: probability})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return alphabet
}

func ReadFiles(path string, extension string) []Alphabet {
	extension = "*." + extension
	files, err := getFiles(path, extension)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {
			return readFile(file)
		}
	}
	return nil
}
