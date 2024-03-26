// https://go.dev/play/

package main

import (
	"testing"
)

func Hello() string {
	return "Hello, world"
}

func TestLastIndex(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
