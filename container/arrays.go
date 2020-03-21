package main

import "fmt"

func modifyArr(arr *[5]int) {
	arr[0] = 233
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 3, 5, 7, 9}
	var grid [2][3]bool
	fmt.Println(arr1, arr2, arr3, grid)
	modifyArr(&arr1)
	fmt.Println(arr1)
}
