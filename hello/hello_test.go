package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("David")
	want := "Hello, David!"

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
