package main

import "fmt"

func find(s []int, t int) (int, int) {
	var t1, t2 int
	var flag = false
	for i, m := range s {
		for j, n := range s[(i + 1):] {
			if m+n == t {
				t1 = i
				t2 = j + (i + 1)
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	return t1, t2
}

type TransInfo int

type Fragment interface {
	Exec(transInfo *TransInfo) error
}

type GetPodAction struct {
}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}

// TODO why
var fragment1 Fragment = new(GetPodAction) // right

// var fragment2 Fragment = GetPodAction // wrong
var fragment3 Fragment = &GetPodAction{} // right
var fragment4 Fragment = GetPodAction{}  // right

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(find(s, 11))
	fmt.Println(find(s, 9))
	fmt.Println(find(s, 3))
}
