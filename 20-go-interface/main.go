package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	Width, Height float64
}

type circle struct {
	Radius float64
}

type sharp interface {
	Area() float64
}

func (c circle) Area() float64 {
	return math.Phi * math.Pow(c.Radius, 2)
}

func (r rectangle) Area() float64 {
	return r.Height * r.Width
}

func main() {
	fmt.Println("=====Ini merupakan contoh Slice di Golang=====")

	r := rectangle{
		Width:  2,
		Height: 4,
	}

	c := circle{5}

	fmt.Println("Area of the rectangle is", getArea(r))
	fmt.Println("Area of the circle is", getArea(c))
}

func getArea(s sharp) float64 {
	return s.Area()
}
