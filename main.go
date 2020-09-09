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
	alphabetNormal := "abcdefghijklmnopqrstuvwxyz" + " "
	alphabetSecret := string(shuffle("abcdefghijklmnopqrstuvwxyz")) + " "
	realQuadgrams := language_tools.ReadQuadramsFromFile("./english_quadgrams.txt")
	alphabetNormalProb := language_tools.ReadFiles("./alphabets", "csv")

	enc := monoalphabetic.Encrypt(message, alphabetNormal, alphabetSecret)
	dec := monoalphabetic.Decrypt(enc, alphabetNormal, alphabetSecret)
	fmt.Println("---------------------------------MONOALPHABETIC---------------------------------")

	fmt.Println(alphabetNormal)
	fmt.Println(alphabetSecret)
	fmt.Println("Encrypted: ", enc)
	fmt.Println("  ")
	fmt.Println("Decrypted: ", dec)
	fmt.Println("Cracking !!!!: ")
	cracked := monoalphabetic.Crack(enc, alphabetNormal, realQuadgrams, alphabetNormalProb)
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
	realQuadgrams := language_tools.ReadQuadramsFromFile("./english_quadgrams.txt")
	alphabetNormalProb := language_tools.ReadFiles("./alphabets", "csv")

	keyWord := "morca"
	fmt.Println("---------------------------------VINEGER---------------------------------")
	enc := vigenere.Encrypt(message, alphabets, keyWord)
	dec := vigenere.Decrypt(enc, alphabets, keyWord)
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Decrypted: ", dec)
	vigenere.Crack(enc, realQuadgrams, alphabetNormalProb, alphabets)
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
	message = message[len(message)/2:]
	// monoalphabeticTest(message)
	if len(os.Args) > 1 && os.Args[1] == "-m" {
		fmt.Println(message)
	} else {
		vigener(message)
	}
	// msg := "PPQCA XQVEKG YBNKMAZU YBNGBAL JON I TSZM JYIM VRAG VOHT VRAU C TKSG DDWUO XITLAZU VAVV RAZ C VKB QP IWPOU"

}
