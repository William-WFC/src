package main

import "fmt"

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func main() {
	q := Queue{1, 2, 3}
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q)
}
