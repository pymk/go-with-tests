package main

import (
	"testing"
)

func assertrCorrectMessage(t testing.TB, got, expected string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function
	// call rather than inside our test helper.
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("world", "")
		expected := "Hello, world!"

		assertrCorrectMessage(t, got, expected)
	})
	t.Run("say 'Hello, world!' when empty strings are supplied", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, world!"

		assertrCorrectMessage(t, got, expected)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Elenoir", french) // french variable is defined in hello.go
		expected := "Bonjour, Elenoir!"

		assertrCorrectMessage(t, got, expected)
	})
}
