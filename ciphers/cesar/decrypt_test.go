package cesar

import "testing"

// "abcdefghijklmnopqrstuvwxyz"

func TestShiftDecrypt(t *testing.T) {
	newChar := shiftDecrypt('m', 3)

	if newChar != 'j' {
		t.Errorf("Error cesar cipher shifting m %c", newChar)
	}

	newChar = shiftDecrypt('a', 3)
	if newChar != 'x' {
		t.Errorf("Error cesar cipher shifting a %d", newChar)
	}

	newChar = shiftDecrypt('b', 3)
	if newChar != 'y' {
		t.Errorf("Error cesar cipher shifting b %d", newChar)
	}

}

func TestDecrypt(t *testing.T) {
	enc := "sri qsvrmrk alir kviksv weqwe asoi"
	dec := "one morning when gregor samsa woke"

	if Decrypt(enc, 4) != dec {
		t.Errorf("Error cesar cipher decrypting %s", enc)
	}
}
