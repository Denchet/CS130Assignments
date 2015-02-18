//covers if/else, switching, functions as args, slices, ranges, loops, break
//and continue
package main

import "fmt"

type functype func(int) int

func double(n int) int {
	return n * 2
}

func printSlice(a []int, b functype) {
	for i := range a {
		switch a[i] {
		case 0:
			fmt.Println(b(a[i]))
		case 1:
			continue
		case 2:
			fmt.Println(b(a[i]))
		case 3:
			continue
		case 4:
			fmt.Println(b(a[i]))
		case 5:
			continue
		case 6:
			break
		}
	}
}

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if intSlice[5] == 5 {
		fmt.Println("the conditional evaluated: true")
	} else {
		fmt.Println("the conditional evaluated: false")
	}

	for i := range intSlice {
		switch intSlice[i] {
		case 0:
			fmt.Println(intSlice[i])
		case 1:
			continue
		case 2:
			fmt.Println(intSlice[i])
		case 3:
			continue
		case 4:
			fmt.Println(intSlice[i])
		case 5:
			continue
		case 6:
			break
		}
	}

	printSlice(intSlice, double)
}
