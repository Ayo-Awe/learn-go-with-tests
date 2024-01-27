package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		msg := Hello("Chris", "English")
		expected := "Hello, Chris"
		assertCorrectMessage(t, msg, expected)
	})

	t.Run("say hello world when name is an empty string", func(t *testing.T) {
		msg := Hello("", "English")
		expected := "Hello, World"
		assertCorrectMessage(t, msg, expected)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		msg := Hello("Marco", "Spanish")
		expected := "Hola, Marco"
		assertCorrectMessage(t, msg, expected)
	})

	t.Run("unrecognized language", func(t *testing.T) {
		msg := Hello("Bola", "Yoruba")
		expected := "Hello, Bola"
		assertCorrectMessage(t, msg, expected)
	})
}

func assertCorrectMessage(t testing.TB, received, expected string) {
	t.Helper()

	if received != expected {
		t.Errorf("got %q, expected %q", received, expected)
	}
}
