package main

import (
	"bufio"
	"fmt"
	"io"
	"learngo/functional/fib"
	"strings"
)

//func fibonacci() intGen {
//	a, b, c := 0, 0, 1
//	return func() int {
//		a = b
//		b = c
//		c = a + b
//		return a
//	}
//}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// TODO: incorrect if p is too small
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func readFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type intGen func() int

func main() {
	var f intGen = fib.Fibonacci()
	readFileContents(f)
}
