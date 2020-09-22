package monoalphabetic

import (
	"testing"
)

func TestTransformEncrypt(t *testing.T) {
	type args struct {
		char           rune
		alphabtNormal  string
		alphabetSecret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Letter a",
			args: args{
				char:           'a',
				alphabtNormal:  "abcdefghijklmnopqrstuvwxyz",
				alphabetSecret: "bacdefghijklmnopqrstuvwxyz",
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformEncrypt(tt.args.char, tt.args.alphabtNormal, tt.args.alphabetSecret); got != tt.want {
				t.Errorf("TransformEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncrypt(t *testing.T) {
	type args struct {
		text           string
		alphabtNormal  string
		alphabetSecret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test string: test case monoalphabetic ",
			args: args{
				text:           "test case monoalphabetic",
				alphabtNormal:  "abcdefghijklmnopqrstuvwxyz",
				alphabetSecret: "baedcfghijmlknopqrstuvwxyz",
			},
			want: "tcst ebsc konoblphbactie",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.text, tt.args.alphabtNormal, tt.args.alphabetSecret); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
