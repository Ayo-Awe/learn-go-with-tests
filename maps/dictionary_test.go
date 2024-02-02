package maps

import "testing"

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "a form of assessment"}
	expected := "a form of assessment"

	definition := dictionary.Search("test")

	if definition != expected {
		t.Errorf("expected %q, but got %q", expected, definition)
	}
}
