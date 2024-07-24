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

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 2},
	}

	for _, shape := range shapes {
		if rect, ok := shape.(Rectangle); ok { // type switch to rectangle
			fmt.Printf("Rectangle area: %.2f\n", rect.Area())
		} else if circle, ok := shape.(Circle); ok { // type switch to circle
			fmt.Printf("Circle area: %.2f\n", circle.Area())
		}
	}
}
