package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(got string,want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q",got,want)
		}
	}

	t.Run("Saying hello to a given name", func(t *testing.T) {
		got:= Hello("Dear", "")
		want:= "Hello, Dear"
		assertCorrectMessage(got, want)
	})

	t.Run("Say hello world when no name given", func(t *testing.T) {
		got := Hello("", "")
		want:= "Hello, world"
		assertCorrectMessage(got, want)
	})

	t.Run("Saying hello to a given name in a recognized language", func(t *testing.T) {
		got:= Hello("Dear", "Spanish")
		want:= "Hola, Dear"
		assertCorrectMessage(got, want)
	})

	t.Run("Saying hello to a given name in french language", func(t *testing.T) {
		got:= Hello("Dear", "French")
		want:= "Bonjour, Dear"
		assertCorrectMessage(got, want)
	})

	t.Run("Saying hello to a given name in a non-recognized language", func(t *testing.T) {
		got:= Hello("Dear", "Japanese")
		want:= "Hello, Dear"
		assertCorrectMessage(got, want)
	})


}
