package arrays

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	checkSum := func(got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got ,want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum collection", func(t *testing.T) {
		numbers := [] int {1,2,3}
		got := sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("SumAll collection", func(t *testing.T) {
		slice1 := [] int {1,2,3,4}
		slice2 := [] int {1,2,3}
		got := sumAll(slice1, slice2)
		want := [] int {10,6}
		checkSum(got, want)
	})

	t.Run("SumTails to collection", func(t *testing.T) {
		slice1 := [] int {1,2,3,4}
		slice2 := [] int {1,2,3}
		got := sumTails(slice1, slice2)
		want := [] int {9,5}
		checkSum(got, want)
	})

	t.Run("SumTails to collection passing empty array", func(t *testing.T) {
		slice1 := [] int {}
		slice2 := [] int {1,2,3}
		got := sumTails(slice1, slice2)
		want := [] int {0,5}
		checkSum(got, want)
	})
}
