package main

import "fmt"

const (
	PI       = 3.14
	Language = "Go"
)

func AreaDiameter(r float32) (float32, float32) {
	var area float32
	var diameter float32

	area = PI * (r * r)
	diameter = PI * (r + r)
	return area, diameter
}

func main() {
	fmt.Println("Please enter the radius of a circle\n")
	var radius float32
	fmt.Scan(&radius)
	fmt.Println(AreaDiameter(radius))
}
