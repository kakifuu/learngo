package main

import (
	"fmt"
	"io/ioutil"
)

const filename = "basic/abc.txt"

func grade(score int) string {
	var grade string
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		grade = "F"
	case score < 80:
		grade = "C"
	case score < 90:
		grade = "B"
	case score <= 100:
		grade = "A"
	}
	return grade
}

func main() {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(1),
		grade(60),
		grade(72),
		grade(81),
		grade(99),
		//grade(101),
	)
}
