package main

import (
	"fmt"

	"github.com/khan745/cipher_cracker/ciphers/cesar"
	"github.com/khan745/cipher_cracker/language_tools"
)

func main() {
	alphabet := language_tools.ReadFiles("./alphabets/", "*.csv")
	fmt.Println(alphabet)
	encrypted := cesar.Encrypt("Ahoj Marek", 3)
	fmt.Println(encrypted)
}
