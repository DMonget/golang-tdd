package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, repeated string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("got %q, but wanted %q", expected, repeated)
		}
	}

	t.Run("Repeating a letter 5 times gives a string of 5 times that letter", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("Repeating a letter x times gives a string of x times that letter", func(t *testing.T) {
		repeated := Repeat("b", 8)
		expected := "bbbbbbbb"
		assertCorrectMessage(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
