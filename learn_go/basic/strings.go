package main

import "fmt"

func main() {
	s := "Yes我爱。。。!" // UTF-8
	fmt.Println(len(s))
	fmt.Printf("%s\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
}
