package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "Bolu"
	buf := bytes.Buffer{}

	Greet(&buf, name)

	expected := "Hello, " + name
	got := buf.String()

	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}
