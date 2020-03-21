package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 233
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	//fmt.Println("arr[2:] ", arr[2:])
	//fmt.Println("arr[2:6] ", arr[2:6])
	//fmt.Println("arr[:6] ", arr[:6])
	//fmt.Println("arr[:] ", arr[:])

	s1, s2 := arr[2:], arr[:6]
	fmt.Printf("Before update\n s1 = %v s2 = %v\n", s1, s2)
	updateSlice(s1)
	updateSlice(s2)
	fmt.Printf("After update\n s1 = %v s2 = %v\n", s1, s2)

	s3 := arr[2:6]
	fmt.Printf("s3 = %v; len(s3) = %d; cap(s3) = %d\n", s3, len(s3), cap(s3))
	fmt.Printf("s3[4:5] = %v\n", s3[4:5])

	ext := [2]int{10, 11}

	s4 := append(s3, ext[:]...)
	s5 := append(s4, 12)

	fmt.Printf("s4 = %v; len(s4) = %d; cap(s4) = %d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5 = %v; len(s5) = %d; cap(s5) = %d\n", s5, len(s5), cap(s5))
}
