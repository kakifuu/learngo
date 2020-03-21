package fib

func Fibonacci() func() int {
	a, b, c := 0, 0, 1
	return func() int {
		a = b
		b = c
		c = a + b
		return a
	}
}
