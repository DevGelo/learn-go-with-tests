// https://go.dev/play/

package main

import (
	"testing"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func TestHello(t *testing.T) {
	got := Hello("Mike")
	want := "Hello, Mike"

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
