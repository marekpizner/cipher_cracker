package main

import (
	"fmt"
	"math/rand"
	"regexp/syntax"
	"time"

	"github.com/khan745/cipher_cracker/ciphers/cesar"
	"github.com/khan745/cipher_cracker/ciphers/transposition/monoalphabetic"
	"github.com/khan745/cipher_cracker/ciphers/vigenere"
)

var message = "One morning when Gregor Samsa woke from troubled dreams he found himself transformed in his bed into a horrible vermin He lay on his armourlike back and if he lifted his head a little he could see his brown belly slightly domed and divided by arches into stiff sections The bedding was hardly able to cover it and seemed ready to slide off any moment His many legs pitifully thin compared with the size of the rest of him waved about helplessly as he looked  What s happened to me  he thought It wasn t a dream His room a proper human room although a little too small lay peacefully between its four familiar walls A collection of textile samples lay spread out on the table  Samsa was a travelling salesman  and above it there hung a picture that he had recently cut out of an illustrated magazine and housed in a nice gilded frame It showed a lady fitted out with a fur hat and fur boa who sat upright raising a heavy fur muff that covered the whole of her lower arm towards the viewer Gregor then turned to look out the window at the dull weather Drops of rain could be heard hitting the pane which made him feel quite sad  How about if I sleep a little bit longer and forget all this nonsense  he thought but that was something he was unable to do because he was used to sleeping on his right and in his present state couldn t get into that position However hard he threw himself onto his right he always rolled back to where he was He must have tried it a hundred times shut his eyes so that he wouldn t have to look at the floundering legs and only stopped when he began to feel a mild dull pain there that he had never felt before  Oh God  he thought  what a strenuous career it is that I ve chosen Travelling day in and day out Doing business like this takes much more effort than doing your own business at home and on top of that there s the curse of travelling worries about making train connections bad and irregular food contact with different people all the time so that you can never get to know anyone or become friendly with them It can all go to Hell  He felt a slight itch up on his belly pushed himself slowly up on his back towards the headboard so that he could lift his head better found where the itch was and saw that it was covered with lots of little white spots which he didn t know what to make of and when he tried to feel the place with one of his legs he drew it quickly back because as soon as he touched it he was overcome by a cold shudder He slid back into his former position  Getting up early all the time  he thought  it makes you stupid You ve got"

func cesarTest() {
	// message := "Ahoj Marek ! :) Cesar"
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

func monoalphabeticTest() {
	// message := "Ahoj Marek ! :) Monoalphabetic"
	alphabetNormal := "! :)0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alphabetSecret := string(shuffle("! :)0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"))

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

func vigener() {
	alphabets := generateVigenerAlphabet()
	message := "Ahoj Marek ! :) Vineger"
	fmt.Println("---------------------------------VINEGER---------------------------------")
	enc := vigenere.Encrypt(message, alphabets, "MORCA")
	dec := vigenere.Decrypt(enc, alphabets, "MORCA")
	fmt.Println("Encrypted: ", enc)
	fmt.Println("Decrypted: ", dec)

}

func foo() {
	word1 := []rune("alpha")
	word2 := rune('吃') // no need for array if for single rune
	word3 := []rune("1234")
	word4 := []rune(" $#$^@#$ ")

	ok := syntax.IsWordChar(word1[0])

	fmt.Printf("%v is a word ? : %v \n", string(word1), ok)

	ok = syntax.IsWordChar(word2)

	fmt.Printf("%v is a word ? : %v \n", string(word2), ok)

	ok = syntax.IsWordChar(word3[0])

	fmt.Printf("%v is a word ? : %v \n", string(word3), ok)

	ok = syntax.IsWordChar(word4[0])

	fmt.Printf("%v is a word ? : %v \n", string(word4), ok)
}

func main() {
	// cesarTest()
	monoalphabeticTest()
	// vigener()
}
