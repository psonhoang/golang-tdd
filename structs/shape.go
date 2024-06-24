package main

import "math"

type Shape interface {
	Area() float64
}

// Rectangle
type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Triangle 
type Triangle struct {
	Base	float64
	Height	float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2.0
}