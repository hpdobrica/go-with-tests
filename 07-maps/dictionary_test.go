package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("retrieves a word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a drill"}

		got := dictionary.Search("test")
		want := "this is just a drill"

		assertStrings(t, got, want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
