package main

import (
	"fmt"
	"learnGo/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://vooglam.com"

func download(r Retriever) string {
	return r.Get("http://vooglam.com")
}

func post(poster Poster) {
	poster.Post("https://vooglam.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked",
	})
	return s.Get(url)
}

var x error = nil

func continueTest() {
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(num)
	}

	str := ""
	fmt.Println(str)

}

func main() {
	var r mock.Retriever
	download(r)
	fmt.Println(session(r))

	s := "魔晶3.0"
	fmt.Println(len(s))
	continueTest()

	fmt.Println(x)

}
