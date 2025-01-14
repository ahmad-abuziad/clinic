package validator

import (
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/assert"
)

func TestValidator(t *testing.T) {

	v := New()

	v.Check(true, "valid_field", "should not be added to errors")
	v.Check(1 == 2, "invalid_field", "should be added to errors")
	v.Check(1 == 3, "invalid_field", "should not be added, because this field already got an error")

	v.AddError("key", "message")

	v.Check(PermittedValue(1, []int{1, 2, 3}...), "valid_permitted_value", "1 is permitted within provided values 1, 2, 3")
	v.Check(PermittedValue(-1, []int{1, 2, 3}...), "invalid_permitted_value", "-1 is not permitted with provided values 1, 2, 3")

	v.Check(Matches("test@example.com", EmailRX), "valid_matches", "test@exmaple.com is matches the email regular expression")
	v.Check(Matches("test@example.", EmailRX), "invalid_matches", "test@exmaple. does not match the email regular expression")

	v.Check(Unique([]string{"a", "b", "c"}), "valid_unique", "the slice values a, b, c are unique")
	v.Check(Unique([]string{"a", "a", "c"}), "invalid_unique", "the slice values a, a, b are not unique")

	assert.Equal(t, v.Valid(), false)
	assert.Equal(t, len(v.Errors), 5)
}
