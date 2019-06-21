package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Table based test
func TestArea(t *testing.T) {
	// anonymous struct
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// struct without named fields (implicit based on Type
		// {Rectangle{12, 6}, 72.0},
		// struct with named fields
		{name: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)

			}
		})
	}
}

// These were the original tests. Basic structure, and repetitive. Refactored to above code.
// checkArea := func(t *testing.T, shape Shape, want float64) {
// 	t.Helper()
// 	got := shape.Area()
//
// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }
//
// t.Run("rectangles", func(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	checkArea(t, rectangle, 72.0)
//
// })
//
// t.Run("circles", func(t *testing.T) {
// 	circle := Circle{10}
// 	checkArea(t, circle, 314.1592653589793)
//
// })
