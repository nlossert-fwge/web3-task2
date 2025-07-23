package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	r := Rectangle{Width: 10, Height: 20}
	c := Circle{Radius: 5}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
}
