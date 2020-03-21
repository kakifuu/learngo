package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

//func tryDefer() {
//	i := 1
//	defer fmt.Println(i)
//	i++
//	defer fmt.Println(i)
//	i++
//	fmt.Println(i)
//	i++
//	defer fmt.Println(i)
//}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	defer w.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(w, f())
	}
}

func main() {
	writeFile("fib.txt")
}
