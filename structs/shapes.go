package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (t Triangle) Area() float64 {
	return t.Height * t.Width / 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Perimeter(r Rectangle) float64 {
	return 2*r.Height + 2*r.Width
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}
