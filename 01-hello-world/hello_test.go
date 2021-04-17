package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("saying hi to people", func(t *testing.T) {
		got := Hello("Dobrica", "")
		want := "Hello, Dobrica"
		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello world when empty string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Dobrica", "Spanish")
		want := "Hola, Dobrica"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Dobrica", "French")
		want := "Bonjour, Dobrica"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in japanese", func(t *testing.T) {
		got := Hello("Dobrica", "Japanese")
		want := "Konnichiwa, Dobrica"

		assertCorrectMessage(t, got, want)
	})

}
