package main

import "fmt"
import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2
	a -= x1
	b := y2
	b -= y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimeter() float64
}
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return c.r * 2 * math.Pi
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l * 2) + (w * 2)
}
func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	mySlice := []int{1, 2, 3, 4, 5}
	mySlice2 := append(mySlice, 6, 7)

	myArray := [7]int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < 7; i++ {
		if myArray[i] == mySlice2[i] {
			fmt.Println(mySlice2[i])
		}
	}

}
