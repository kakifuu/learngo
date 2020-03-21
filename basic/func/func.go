package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	q, r = a/b, a%b
	return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(number ...int) int {
	sum := 0
	for _, v := range number {
		sum += v
	}
	return sum
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	if result, err := eval(14, 6, "?"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	q, r := div(12, 5)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
