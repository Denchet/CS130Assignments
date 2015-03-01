package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	x := 5
	y := 10
	swap(&x, &y)
	fmt.Println(x, y)
}
