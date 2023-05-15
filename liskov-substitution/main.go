package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 5, Height: 4},
		Square{SideLength: 3},
	}

	totalArea := CalculateTotalArea(shapes)
	fmt.Println("Total Area:", totalArea)
}
