package main

import (
	"bufio"
	"fmt"
	"learnGo/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	// panic("error occurred")
	fmt.Println(4)
}

func Error(err error) string {
	return "custom err"
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// err = errors.New("this is a custom err")
	if err != nil {
		// fmt.Println(err.Error())
		fmt.Println(Error(err))
		// if pathError, ok := err.(*os.PathError); !ok {
		// 	panic(err)
		// } else {
		// 	fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		// }
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
	var x = []int{1, 5: 4, 6}
	fmt.Println(x)
}
