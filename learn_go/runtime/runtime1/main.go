package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 5; i++ {
		// 切一下，再次分配任务
		fmt.Println("hello")
		runtime.Gosched()
	}
}
