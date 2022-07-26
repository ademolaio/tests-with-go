package main

import "testing"

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
