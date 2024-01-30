package structsmethodsinterfaces

import "testing"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		perimeter := shape.Perimeter()

		if perimeter != expected {
			t.Errorf("expected %g, but got %g", expected, perimeter)
		}
	}

	t.Run("test rectangle", func(t *testing.T) {
		rectangle := Rectangle{20, 10}
		expected := 60.0
		checkPerimeter(t, rectangle, expected)
	})

	t.Run("test circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 62.83185307179586
		checkPerimeter(t, circle, expected)
	})

}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"Rectangle", Rectangle{20, 10}, 200},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{2, 8}, 8},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			area := tt.shape.Area()
			if area != tt.expected {
				t.Errorf("expected %g, but got %g", tt.expected, area)
			}
		})
	}
}
