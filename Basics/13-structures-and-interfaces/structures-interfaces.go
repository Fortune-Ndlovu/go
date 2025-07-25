package main
import (
	"fmt"
	"math"
)

func main () {
	rect := Rectangle{50,60}
	circ := Circle{7}
	fmt.Println("Area of rectangle is ", getArea(rect))
	fmt.Println("Radius of cicle is ", getArea(circ))
// 	Area of rectangle is  3000
// Radius of cicle is  153.93804002589985
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func ( r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius,2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

