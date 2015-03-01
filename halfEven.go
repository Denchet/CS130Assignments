package main

import "fmt"

func halfEven(x int) (int, bool) {
	r := x / 2
	if x%2 == 0 {
		return r, true
	} else {
		return r, false
	}
}

func main() {
	var n int
	var b bool

	n, b = halfEven(1)
	fmt.Println(n, b)
}
