package main

import "fmt"

func main() {
	m1 := map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	}
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m1, m2, m3)

	fmt.Println("traversing map")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("getting vaues")
	cousrseName, ok := m1["course"]
	fmt.Println(cousrseName, ok)

	if cousrseName, ok := m1["course"]; ok {
		fmt.Println("if", cousrseName)
	} else {
		fmt.Println("key does not exiting")
	}

	fmt.Println("delete values")
	name, ok := m1["name"]
	fmt.Println(name, ok)
	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Println(name, ok)
}
