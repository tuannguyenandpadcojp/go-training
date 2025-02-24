package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area of the %T %+v is %0.2f\n", s, s, s.Area())
	fmt.Printf("Perimeter of the %T %+v is %0.2f\n", s, s, s.Perimeter())
}

type Circle struct {
	R float64
}

func (c Circle) Area() float64 {
	return c.R * c.R * math.Pi // π * R²
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.R // 2 * π * R
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func main() {
	// Circle
	c := Circle{R: 3}
	printShapeInfo(c)

	// Rectangle
	r := Rectangle{Width: 3, Height: 4}
	printShapeInfo(r)
}
