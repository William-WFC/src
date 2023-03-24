package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
)
var cc complex64 = 12

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func varialbleShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s, aa, ss)
}

func euler() {
	dd := 3 + 4i
	fmt.Println(dd, cmplx.Abs(dd))

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)

}

func triangle() {
	var a, b int = 3, 4

	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int

	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
}

func enums() {
	const (
		cjj = iota
		_
		pathon
		abc
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	const (
		aa = 10 * iota
		bb
		cc
		dd
	)
	fmt.Println(cjj, pathon, abc, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(math.Pow(2, 10))
	fmt.Println(aa, bb, cc, dd)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	varialbleShorter()
	euler()
	triangle()
	consts()
	enums()
}
