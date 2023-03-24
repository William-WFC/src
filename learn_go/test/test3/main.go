package main

import "fmt"

// 设计以高度 n 为参数的一个函数，实现等腰三角形的打印；

func triggle(n int) {
	if n < 1 {
		fmt.Println("参数n必须大于等于1")
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d  ", i)
		}
		fmt.Println()
	}

	fmt.Printf("高度为%d的等腰三角形打印完毕\n", n)
}

func main() {
	triggle(1)
	triggle(3)
	triggle(10)

}
