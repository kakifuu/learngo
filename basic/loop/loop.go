package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	readFileContents(file)
}

func readFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func forever() {
//	for {
//		fmt.Println("Hi")
//	}
//}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0),
	)
	readFile("basic/abc.txt")
	//forever()
	s := `Why think separately of 
this life and the next
when one is born from the last`
	readFileContents(strings.NewReader(s))
}
