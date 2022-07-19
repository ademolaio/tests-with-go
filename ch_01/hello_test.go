package main

import "testing"

func TestHello(t *testing.T)  {
	t.Run("Saying hello to people", func(t *testing.T) {

		got := Hello("Chris")
		want := "Hello, Chris"
	
		if got != want {
			t.Errorf("Got %q, but want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func (t *testing.T)  {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q, but want %q", got, want)
		}
	})

}