package main

import "fmt"

type Additionmessage struct {
	message string
	input1  int
	input2  int
	result  int
}

func main() {
	var m = Additionmessage{}
	m.message = "Addition example: "

	fmt.Println("Enter two numbers to be added, seperated by a space")

	fmt.Scan(&m.input1)
	fmt.Scan(&m.input2)
	m.result = m.input1 + m.input2
	fmt.Println(m.message)
	fmt.Println(m.input1, "+", m.input2, "=", m.result)
}
