package cesar

import "testing"

// "abcdefghijklmnopqrstuvwxyz"

func TestShiftEncrypt(t *testing.T) {
	newChar := shiftEncrypt('m', 3)

	if newChar != 'p' {
		t.Errorf("Error cesar cipher shifting m %c", newChar)
	}

	newChar = shiftEncrypt('x', 3)
	if newChar != 'a' {
		t.Errorf("Error cesar cipher shifting x %c", newChar)
	}

	newChar = shiftEncrypt('y', 3)
	if newChar != 'b' {
		t.Errorf("Error cesar cipher shifting b %c", newChar)
	}
}

func TestEncrypt(t *testing.T) {
	enc := "sri qsvrmrk alir kviksv weqwe asoi"
	dec := "one morning when gregor samsa woke"

	if Encrypt(dec, 4) != enc {
		t.Errorf("Error cesar cipher decrypting %s", dec)
	}
}
