package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v; want: %v", got, want)
	}
}

func StringContains(t *testing.T, got, wantSubstring string) {
	t.Helper()

	if !strings.Contains(got, wantSubstring) {
		t.Errorf("got: %v; wanted to contain: %v", got, wantSubstring)
	}
}

func NilError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got: %v; want: nil", got)
	}
}

func NonNilError(t *testing.T, got error) {
	t.Helper()

	if got == nil {
		t.Errorf("got: nil; want: error")
	}
}
