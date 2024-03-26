// https://go.dev/play/

package main

import (
	"testing"
)

func Hello(name string) string {
	return "Hello, " + name
}

func TestHello(t *testing.T) {
	got := Hello("Mike")
	want := "Hello, Mike"

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
