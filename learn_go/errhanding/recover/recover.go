package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err.Error())
		} else {
			panic(fmt.Sprintf("I dont know what to do: %v", r))
			// panic(r)
		}
	}()
	// panic(errors.New("this is a custom error"))
	// var b = 0
	// var a = 4 / b
	// fmt.Println(a)

	panic(13)

}

func main() {
	tryRecover()
}
