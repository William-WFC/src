package main

import "fmt"

func judge(v int) bool {
	flag := false
	var x int
	for i, j := 1, 1; j <= v; {
		if v == j {
			flag = true
			break
		}
		x = j
		j += i
		i = x
	}
	return flag
}

func main() {
	fmt.Println(judge(28656))
	fmt.Println(judge(28657))
	fmt.Println(judge(28658))
}
