package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jane")
	want := "Hello, Jane"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
