package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"

	"github.com/khan745/cipher_cracker/ciphers/cesar"
)

func cesarTest() {
	message := "Ahoj Marek ! :) Cesar"
	enc := cesar.Encrypt(message, 4)
	dec := cesar.Decrypt(enc, 4)
	fmt.Println("---------------------------------CESAR---------------------------------")
	fmt.Println("Encrypted: %s", enc)
	fmt.Println("Decrypted: %s", dec)
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

func monoalphabeticTest() {
	message := "Ahoj Marek ! :) Monoalphabetic"
	alphabetNormal := "! :) 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alphabetSecret := string(shuffle("! :) 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"))

	enc := monoalphabetic.Encrypt(message, alphabetNormal, alphabetSecret)
	dec := monoalphabetic.Decrypt(enc, alphabetNormal, alphabetSecret)
	fmt.Println("---------------------------------MONOALPHABETIC---------------------------------")
	fmt.Println("Encrypted: %s", enc)
	fmt.Println("Decrypted: %s", dec)

}

func generateVigenerAlphabet() []string {
	alphabet := "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}"
	alphabetLength := len(alphabet)
	var alphabets []string

	for i := 0; i < alphabetLength; i++ {
		newAlphabet := alphabet[i:] + alphabet[:i]
		alphabets = append(alphabets, newAlphabet)
	}
	return alphabets
}

func main() {
	cesarTest()
	monoalphabeticTest()

	alphabets := generateVigenerAlphabet()
	for _, x := range alphabets {
		fmt.Println(x)
	}
}
