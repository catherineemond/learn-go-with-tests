package iteration

import (
	"testing"
)

func TestReverse(t *testing.T) {
	reversed := Reverse("hello")
	expected := "olleh"

	if reversed != expected {
		t.Errorf("expected %q but got %q", expected, reversed)
	}
}
