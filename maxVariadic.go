package main

import "fmt"

func max(x ...int) int {
	max := 0
	for _, value := range x {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println(max(1, 2, 17, 4, 5))
}
