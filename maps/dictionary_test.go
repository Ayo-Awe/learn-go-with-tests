package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "a form of assessment"}
		expected := "a form of assessment"

		definition, _ := dictionary.Search("test")
		assertStrings(t, definition, expected)
	})

	t.Run("word not found", func(t *testing.T) {
		dictionary := Dictionary{"test": "a form of assessment"}

		_, err := dictionary.Search("rust")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, str, expected string) {
	t.Helper()

	if str != expected {
		t.Errorf("expected %q, but got %q", expected, str)
	}
}

func assertError(t testing.TB, err, expected error) {
	t.Helper()

	if err != expected {
		t.Errorf("expected err: %v, but got err: %v", expected, err)
	}
}
