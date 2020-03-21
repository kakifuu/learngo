package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界!😄"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("%v\n", []byte(s))
	for i, ch := range s {
		fmt.Printf("(%d, %X)", i, ch)
	}

	fmt.Println("\nRune Count: ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
}
