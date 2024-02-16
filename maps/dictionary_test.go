package maps

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "a form of assessment"}
		expected := "a form of assessment"

		definition, err := dictionary.Search("test")

		assertNoError(t, err)
		assertStrings(t, definition, expected)
	})

	t.Run("word not found", func(t *testing.T) {
		dictionary := Dictionary{"test": "a form of assessment"}

		_, err := dictionary.Search("rust")

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("add word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("shoes", "foot protection")

		definition, err := dictionary.Search("shoes")

		assertNoError(t, err)
		assertStrings(t, definition, "foot protection")
	})

	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{}
		word, definition := "shoes", "foot protection"
		dictionary.Add(word, definition)

		err := dictionary.Add(word, "foot wear")
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should successfully update word", func(t *testing.T) {
		dict := Dictionary{}
		word := "code"
		definition := "foo bar"

		err := dict.Add(word, definition)
		assertNoError(t, err)

		newDefinition := "computer instructions"

		err = dict.Update(word, newDefinition)
		assertNoError(t, err)

		updatedDefinition, err := dict.Search(word)
		assertNoError(t, err)
		assertStrings(t, updatedDefinition, newDefinition)
	})
}

func TestDelete(t *testing.T) {
	word := "books"
	dict := Dictionary{}

	err := dict.Add(word, "brain food")
	assertNoError(t, err)

	dict.Delete(word)
	assertNoError(t, err)

	_, err = dict.Search(word)
	assertError(t, err, ErrNotFound)
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

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected nil but got an error")
	}
}
