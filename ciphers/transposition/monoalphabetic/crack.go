package monoalphabetic

import "fmt"

func check(s string) map[rune]uint {
	m := make(map[rune]uint, len(s))
	for _, r := range s {
		m[r]++
	}
	return m
}

func Crack(text string) {
	frequenties := check(text)
	for i, x := range frequenties {
		fmt.Println(string(i), x)
	}
}
