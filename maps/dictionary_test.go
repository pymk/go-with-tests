package main

import "testing"

const (
	test = "test"
	definition = "this is just a test"
	newDefinition = "new definition"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{test: definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(test)
		want := definition
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})

	t.Run("existing word", func(t *testing.T) {
		word := test
		definition := definition
		newDefinition := newDefinition
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := test
		definition := definition
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := test
	definition := definition
	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
}

func TestUpdate(t *testing.T) {
	word := test
	definition := definition
	dictionary := Dictionary{word: definition}
	newDefinition := newDefinition

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}

func TestDelete(t *testing.T) {
	word := test
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

// Helper functions -----------------------------------------------
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if definition != got {
		t.Errorf("got %q, wanted %q", got, definition)
	}
}