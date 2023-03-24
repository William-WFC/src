package main

import (
	"bufio"
	"fmt"
	"io"
	"learnGo/functional/fib"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
	// s1 := fmt.Sprint("枯藤")
	// name := "枯藤"
	// age := 18
	// s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	// s3 := fmt.Sprintln("枯藤")
	// fmt.Println(s1, s2, s3)

	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
}
