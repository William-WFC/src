package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
		// runtime.Gosched()

	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
		// runtime.Gosched()
	}
}

func main() {
	runtime.GOMAXPROCS(20)
	go a()
	go b()
	var ch chan int
	fmt.Println(ch) // <nil>
	time.Sleep(time.Second)

}
