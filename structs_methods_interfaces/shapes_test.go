package main

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := Perimeter(r)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		s       Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{Width: 12, Height: 6}, 72.0},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{"Triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, actually has area of %g", tt.s, got, tt.hasArea)
			}
		})
	}
}
