// https://go.dev/play/

package main

import (
	"testing"
)

const french = "French"
const spanish = "Spanish"
const italian = "Italian"

const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const italianHelloPrefix = "Buongiorno, "
const englishHelloPrefix = "Hello, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	}

	return prefix + name
}

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello, Mike"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Luis", spanish)
		want := "Hola, Luis"

		assertCorrectMessage(t, got, want)
	})
	t.Run("In French", func(t *testing.T) {
		got := Hello("Coline", french)
		want := "Bonjour, Coline"

		assertCorrectMessage(t, got, want)
	})
	t.Run("In Italian", func(t *testing.T) {
		got := Hello("Paulo", italian)
		want := "Buongiorno, Paulo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q got %q", want, got)
	}
}
