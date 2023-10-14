package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	// The t.Helper() tells the test suite that this method is a helper.
	// This is so when the test fails, the failed line number reported will
	// be in the function call rather than inside the test helper.
	t.Helper()

	if got != want {
		t.Errorf("Actual: %q | Expected: %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Rustin", "English")
		want := "Hello, Rustin"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Japanese", func(t *testing.T) {
		got := Hello("Rustin", "Japanese")
		want := "こんにちは, Rustin"

		assertCorrectMessage(t, got, want)
	})
}
