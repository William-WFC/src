package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("slice=%v len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("create slice")
	var s []int // zero value for slice is nil
	fmt.Println(s, s == nil)
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2, s1)
	printSlice(s2)
	printSlice(s1)
	fmt.Println("deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	// s5 = s2[4:]
	// s6 = s2[:3]
	// s7 = copy(s5, s6)
	printSlice(s2)

	fmt.Println("popping from front")
	front := s2[0]
	s2 = s2[1:]

	printSlice(s2)
	fmt.Println("popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(front, tail)
	printSlice(s2)

}
