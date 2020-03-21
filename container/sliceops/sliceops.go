package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("slice = %d, len = %d, cap = %d\n", s, len(s), cap(s))
}

func deleteElement(s []int, idx int) ([]int, error) {
	switch {
	case idx == 0:
		_, s = popHead(s)
		return s, nil
	case idx == len(s)-1:
		_, s = popTail(s)
		return s, nil
	case idx > 0 && idx < len(s)-1:
		s = append(s[:idx], s[idx+1:]...)
		return s, nil
	default:
		return nil, fmt.Errorf("index out of range [%d] with length %d", idx, len(s))
	}
}

func popHead(s []int) (int, []int) {
	head := s[0]
	s = s[1:]
	return head, s
}

func popTail(s []int) (int, []int) {
	tail := s[len(s)-1]
	s = s[:len(s)-1]
	return tail, s
}

func main() {
	var s []int
	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, 2*i+1)
	}

	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	printSlice(s1)

	s2 := make([]int, 17)
	printSlice(s2)

	s3 := make([]int, 10, 16)
	printSlice(s3)

	copy(s2, s1)
	printSlice(s2)

	// the len & cap of dest slice will not change
	s4 := []int{1, 2}
	printSlice(s4)
	copy(s4, s1)
	printSlice(s4)

	fmt.Printf("Before | s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	if s1, err := deleteElement(s1, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("After | s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	}

}
