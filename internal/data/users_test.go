package data

import (
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/assert"
)

func TestPassword(t *testing.T) {
	var pass password

	pass.Set("pa55word")

	match, err := pass.Matches("pa55word")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, match, true)
	assert.Equal(t, len(pass.hash), 60)
	assert.Equal(t, *pass.plaintext, "pa55word")
}
