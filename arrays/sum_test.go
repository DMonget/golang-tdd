package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Adding the entries of an array containing the numbers 1 to 5 should give 15.", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := (Sum(numbers))
		want := 15

		if got != want {
			t.Errorf("got %d, but wanted %d for %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func TestSumAll(t *testing.T) {

	t.Run("Adding the entries of a couple of arrays should give an array containing a couple of entries.", func(t *testing.T) {
		got := (SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3}))
		want := []int{15, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	fmt.Println(sums)
	// Output: [15 6]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum the entries of the given arrays, leaving out the first entry of each.", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("Sum the entries of the given empty array, without failing.", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
