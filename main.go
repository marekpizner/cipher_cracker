package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/khan745/cipher_cracker/languagetools"

	"github.com/khan745/cipher_cracker/ciphers/cesar"
	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"
	"github.com/khan745/cipher_cracker/ciphers/vigenere"
)

func cesarTest(message string) {
	enc := cesar.Encrypt(message, 4)
	dec := cesar.Decrypt(enc, 4)
	fmt.Println("---------------------------------CESAR---------------------------------")
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Decrypted: ", dec)
	// cesar.Crack(enc)
}

func shuffle(src string) string {
	final := []byte(src)
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return string(final)
}

func monoalphabeticTest(message string) {
	alphabetNormal := "abcdefghijklmnopqrstuvwxyz" + " "
	alphabetSecret := string(shuffle("abcdefghijklmnopqrstuvwxyz")) + " "
	realQuadgrams := languagetools.ReadQuadramsFromFile("./english_quadgrams.txt")
	alphabetNormalProb := languagetools.ReadFiles("./alphabets", "csv")

	enc := monoalphabetic.Encrypt(message, alphabetNormal, alphabetSecret)
	dec := monoalphabetic.Decrypt(enc, alphabetNormal, alphabetSecret)
	fmt.Println("---------------------------------MONOALPHABETIC---------------------------------")

	fmt.Println(alphabetNormal)
	fmt.Println(alphabetSecret)
	fmt.Println("Encrypted: ", enc)
	fmt.Println("  ")
	fmt.Println("Decrypted: ", dec)
	fmt.Println("Cracking !!!!: ")
	cracked, _ := monoalphabetic.Crack(enc, alphabetNormal, realQuadgrams, alphabetNormalProb)
	fmt.Println("Decrypted: ", cracked)
}

func generateVigenerAlphabet() []string {
	// alphabet := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}"
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	alphabetLength := len(alphabet)
	var alphabets []string

	for i := 0; i < alphabetLength; i++ {
		newAlphabet := alphabet[i:] + alphabet[:i]
		alphabets = append(alphabets, newAlphabet)
	}
	return alphabets
}

func vigener(message string) {
	alphabets := generateVigenerAlphabet()
	realQuadgrams := languagetools.ReadQuadramsFromFile("./english_quadgrams.txt")
	alphabetNormalProb := languagetools.ReadFiles("./alphabets", "csv")

	keyWord := "python"
	fmt.Println("---------------------------------VINEGER---------------------------------")
	enc := vigenere.Encrypt(message, alphabets, keyWord)
	dec := vigenere.Decrypt(enc, alphabets, keyWord)
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Decrypted: ", dec)
	vigenere.Crack(enc, realQuadgrams, alphabetNormalProb, alphabets)
}

func foo() {
	a := languagetools.ReadFiles("./alphabets", "csv")
	for x, y := range a {
		fmt.Println(x, y)
	}
}

func clearMessage(str string) string {
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, "!", "", -1)
	str = strings.Replace(str, "?", "", -1)
	str = strings.Replace(str, ":", "", -1)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, "\"", "", -1)
	str = strings.Replace(str, "'", "", -1)
	str = strings.Replace(str, "'", "", -1)
	str = strings.Replace(str, ";", "", -1)
	str = strings.Replace(str, "—", "", -1)
	return str
}

func readMessage() (string, error) {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	str := string(b)
	str = clearMessage(str)
	return str, err
}
func main() {
	message, _ := readMessage()
	message = strings.ToLower(message)
	message = message[:len(message)/2]

	if len(os.Args) > 1 && os.Args[1] == "-m" {
		fmt.Println(message)
	} else {
		vigener(message)
	}

}
