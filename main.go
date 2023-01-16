package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	sep := strings.Repeat("-", 10)

	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)

	fmt.Println(sep, "Reverse", sep)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	rev = BetterReverse(input)
	doubleRev = BetterReverse(rev)
	fmt.Println(sep, "BetterReverse", sep)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	rev, err := CorrectReverse(input)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	doubleRev, err = CorrectReverse(rev)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(sep, "CorrectReverse", sep)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func BetterReverse(s string) string {
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("runes: %q\n", s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func CorrectReverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
