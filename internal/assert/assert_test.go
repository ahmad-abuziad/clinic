package assert

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Run("a equal a", func(t *testing.T) {
		Equal(t, "a", "a")
	})

	t.Run("a not equal b", func(t *testing.T) {
		tInstance := &testing.T{}
		Equal(tInstance, "a", "b")

		if !tInstance.Failed() {
			t.Error("a not equal b, should fail")
		}
	})
}

func TestStringContains(t *testing.T) {
	t.Run("text contains te", func(t *testing.T) {
		StringContains(t, "text", "te")
	})

	t.Run("text does not contain fe", func(t *testing.T) {
		tInstance := &testing.T{}
		StringContains(tInstance, "text", "fe")

		if !tInstance.Failed() {
			t.Error("text does not contain fe, should fail")
		}
	})
}

func TestNilError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		NilError(t, nil)
	})

	t.Run("non nill error", func(t *testing.T) {
		tInstance := &testing.T{}
		NilError(tInstance, errors.New("error message"))

		if !tInstance.Failed() {
			t.Error("non nil error, should fail")
		}
	})
}

func TestNonNilError(t *testing.T) {
	t.Run("non nil error", func(t *testing.T) {
		NonNilError(t, errors.New("error message"))
	})

	t.Run("nill error", func(t *testing.T) {
		tInstance := &testing.T{}
		NonNilError(tInstance, nil)

		if !tInstance.Failed() {
			t.Error("nil error, should fail")
		}
	})
}
