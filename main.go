package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/khan745/cipher_cracker/language_tools"

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
	alphabetNormal := "abcdefghijklmnopqrstuvwxyz"
	alphabetSecret := string(shuffle("abcdefghijklmnopqrstuvwxyz"))

	enc := monoalphabetic.Encrypt(message, alphabetNormal, alphabetSecret)
	dec := monoalphabetic.Decrypt(enc, alphabetNormal, alphabetSecret)
	fmt.Println("---------------------------------MONOALPHABETIC---------------------------------")
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Decrypted: ", dec)
	monoalphabetic.Crack(enc)
}

func generateVigenerAlphabet() []string {
	alphabet := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}"
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
	fmt.Println("---------------------------------VINEGER---------------------------------")
	enc := vigenere.Encrypt(message, alphabets, "MORCA")
	dec := vigenere.Decrypt(enc, alphabets, "MORCA")
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Decrypted: ", dec)

}

func foo() {
	a := language_tools.ReadFiles("./alphabets", "csv")
	for x, y := range a {
		fmt.Println(x, y)
	}
}

func readMessage() (string, error) {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	newString := string(b)
	newString = strings.Replace(newString, ".", "", -1)
	newString = strings.Replace(newString, ",", "", -1)
	newString = strings.Replace(newString, "!", "", -1)
	newString = strings.Replace(newString, "?", "", -1)
	newString = strings.Replace(newString, ":", "", -1)
	newString = strings.Replace(newString, "-", "", -1)
	newString = strings.Replace(newString, "\"", "", -1)
	newString = strings.Replace(newString, "'", "", -1)
	newString = strings.Replace(newString, "'", "", -1)
	newString = strings.Replace(newString, ";", "", -1)

	return newString, err
}
func main() {
	message, _ := readMessage()
	message = strings.ToLower(message)
	// cesarTest(message)
	monoalphabeticTest(message)
	// vigener()
	fmt.Println("asd")
}
