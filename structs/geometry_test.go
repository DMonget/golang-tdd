package structs

import "testing"

func AssertCorrectMessage(t *testing.T, got, want float64) {
	t.Helper()
	if want != got {
		t.Errorf("got %.2f, but wanted %.2f", got, want)
	}
}
func TestPerimiter(t *testing.T) {
	t.Run("The perimiter of a 10.0 * 10.0 rectangle should be 40.0", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimiter()
		want := 40.0

		AssertCorrectMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "The area of a 10.0 height * 10.0 width rectangle should be 100.0", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "The area of a 10.0 radius circle should be 314.1592653589793", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "The area of a 12.0 base * 6.0 height triangle should be 36.0 ", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, entry := range areaTests {
		// using entry.name from the case to use it as the `t.Run` test name
		t.Run(entry.name, func(t *testing.T) {
			got := entry.shape.Area()
			if got != entry.want {
				t.Errorf("%#v got %.2f want %.2f", entry, got, entry.want)
			}
		})
	}

}
