package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32.0) * (5.0 / 9.0)
	fmt.Println(output)
}
