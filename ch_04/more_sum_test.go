package main

import (
	"reflect"
	"testing"
)

func TestMoreSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := MoreSum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := MoreSum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d want %d given, %v", got, want, numbers)
		}
	})
}




func TestSumAllTails(t *testing.T)  {

	t.Run("Make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int {0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but want %v", got, want)
		}
	})

}