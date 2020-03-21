package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue() {
	var (
		a int
		s string
	)
	fmt.Printf("%d %q\n", a, s)
}
func variableInitValue() {
	var (
		a, b = 1, 2
		s    = "abc"
	)
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, s = 1, 2, "hi"
	fmt.Println(a, b, s)
}

func variableShortDeclaration() {
	a, b, s := 1, 2, "hello"
	fmt.Println(a, b, s)
}

func euler() {
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int
	a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func constants() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enum() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello, 世界")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShortDeclaration()

	euler()
	triangle()

	constants()
	enum()
}
