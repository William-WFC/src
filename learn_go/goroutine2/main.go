package main

import (
	"fmt"
	"sync"
)

func hello() {
	fmt.Println("hello goroutine")
}

var wg sync.WaitGroup

func hello2(i int) {
	defer wg.Done()
	fmt.Println("hello goroutine2", i)
}

func main() {
	go hello()
	fmt.Println("hello main")
	// time.Sleep(time.Millisecond)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello2(i)
	}
	wg.Wait()
}
