package main

import (
	"math"
	"testing"
)

func TestShapeInterfaces(t *testing.T) {
	type want struct {
		area      float64
		perimeter float64
	}
	tests := []struct {
		name  string
		shape Shape
		want  want
	}{
		{
			name:  "Circle",
			shape: Circle{R: 5},
			want: want{
				area:      math.Pi * math.Pow(5, 2), // π * R²
				perimeter: 2 * math.Pi * 5,          // 2 * π * R
			},
		},
		{
			name:  "Rectangle",
			shape: Rectangle{Width: 3, Height: 4},
			want: want{
				area:      12, // width * height
				perimeter: 14, // (width + height) * 2
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shape.Area(); got != tt.want.area {
				t.Errorf("%T.Area() = got %0.2f, want %0.2f", tt.shape, got, tt.want.area)
			}
			if got := tt.shape.Perimeter(); got != tt.want.perimeter {
				t.Errorf("%T.Perimeter() = got %0.2f, want %0.2f", tt.shape, got, tt.want.perimeter)
			}
		})
	}
}
