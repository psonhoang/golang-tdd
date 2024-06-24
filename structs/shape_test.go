package main

import "testing"

func assert(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	assert(t, got, want)
}

func TestArea(t *testing.T) {

	// Table-driven tests
	areaTests := []struct{
		name	string
		shape	Shape
		want	float64
	} {
		{"Rectangle", Rectangle{Width: 12, Height: 6}, 72.0},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}