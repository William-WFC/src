package main

import "fmt"

// 12. 下列定义，会报错的是：A (逢超
// A. var x = []int{1, 5: 4, 6} // 使用下标定义值
// B. type person struct{ name string }
// C. ages := make(map[int][]string, 10, 5) // 多传了一个容量的参数
// D. var i interface{}

func main() {
	var x = [...]int{1, 5: 6}
	// var y = [7]int{1,,,,,, 7}
	fmt.Println(x)
	ages := make(map[int][]string, 10)
	ages[1] = []string{}
	fmt.Println(ages)
	fmt.Println([]int{} == nil)

}
