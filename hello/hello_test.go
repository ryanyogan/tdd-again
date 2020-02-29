package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q and wanted %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Ryan", "")
		want := "Hello, Ryan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "SP")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Yogan", "FR")
		want := "Bonjour, Yogan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, World when arguments are empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
