package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}
	arr3 := [...]int{4, 5, 6, 7, 9}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	fmt.Println("printArray(arr1)")
	printArray(&arr1)
	fmt.Println("printArray(arr3)")
	printArray(&arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(&arr1, &arr3)
}
