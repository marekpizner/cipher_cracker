package languagetools

import (
	"testing"
)

func TestCalculateProbability(t *testing.T) {
	str := "AAAAaaabc"
	res := CalculateProbability(str)
	if res['A'] != 4 {
		t.Errorf("Error calculate character probability in string")
	}
	if res['a'] != 3 {
		t.Errorf("Error calculate character probability in string")
	}
	if res['b'] != 1 {
		t.Errorf("Error calculate character probability in string")
	}
}

func TestFindIndexOfString(t *testing.T) {
	str := "HiThisIsX!"
	res := FindIndexOfString(str, 'X')

	if res != 8 {
		t.Errorf("Error: TestFindIndexOfString can not find index of character %s in string %s", str, "X")
	}
}

func TestSwapCharactersInAlphabet(t *testing.T) {
	str := "abcdefghijklmnopqrstuvwxyz"
	char1 := 'a'
	char2 := 'z'

	res := SwapCharactersInAlphabet(str, char1, char2)
	if res != "zbcdefghijklmnopqrstuvwxya" {
		t.Errorf("Error: TestSwapCharactersInAlphabet can not swap characters")

	}
}

func TestSortAlphabet(t *testing.T) {
	t.Parallel()
	alphabetNormal := "abcd"
	alphabetSecret := "jocz"
	alphabetProb := "cdba"

	an, as := sortAlphabet(alphabetProb, alphabetSecret, alphabetNormal)

	if an != alphabetNormal {
		t.Errorf("Wrong order of normal alphabet")
	}

	if as != "zcjo" {
		t.Error("Wrong order of secret alphabet")
	}
}
