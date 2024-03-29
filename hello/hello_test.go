package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}

	t.Run("Saying hello to people.", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people in the chosen language.", func(t *testing.T) {
		got := Hello("David", "french")
		want := "Bonjour, David!"
		assertCorrectMessage(t, got, want)
	})

}
