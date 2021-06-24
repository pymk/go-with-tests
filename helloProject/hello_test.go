package integers

import "testing"

func TestHello(t *testing.T) {

	// Assertion function
	assertrCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Jon", "")
		want := "Hello, Jon!"
		assertrCorrectMessage(t, got, want)
	})

	t.Run("say hello in Farsi", func(t *testing.T) {
		got := Hello("Payam", "Farsi")
		want := "Salam, Payam!"
		assertrCorrectMessage(t, got, want)
	})

	t.Run("say hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertrCorrectMessage(t, got, want)
	})
}
