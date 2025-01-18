package data

import (
	"testing"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/assert"
	"github.com/ahmad-abuziad/clinic/internal/validator"
)

func TestValidatePatient(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		v := validator.New()
		patient := &Patient{
			FirstName:   "Ahmad",
			LastName:    "Abuziad",
			Gender:      "M",
			DateOfBirth: time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC),
		}

		ValidatePatient(v, patient)

		assert.Equal(t, v.Valid(), true)
		assert.Equal(t, len(v.Errors), 0)
	})

	t.Run("default values", func(t *testing.T) {
		v := validator.New()
		patient := &Patient{}

		ValidatePatient(v, patient)

		assert.Equal(t, v.Errors["first_name"], "must be provided")
		assert.Equal(t, v.Errors["last_name"], "must be provided")
		assert.Equal(t, v.Errors["gender"], "must be provided")
		assert.Equal(t, v.Errors["date_of_birth"], "must not be before 1900-01-01")
		assert.Equal(t, v.Valid(), false)
		assert.Equal(t, len(v.Errors), 4)
	})

	t.Run("beyond max values", func(t *testing.T) {
		v := validator.New()
		patient := &Patient{
			FirstName:   "123456789012345678901234567890123456789012345678901",
			LastName:    "123456789012345678901234567890123456789012345678901",
			Gender:      "P",
			DateOfBirth: time.Now().Add(1 * time.Hour),
		}

		ValidatePatient(v, patient)

		assert.Equal(t, v.Errors["first_name"], "must not be more than 50 bytes long")
		assert.Equal(t, v.Errors["last_name"], "must not be more than 50 bytes long")
		assert.Equal(t, v.Errors["gender"], "must be M or F only")
		assert.Equal(t, v.Errors["date_of_birth"], "must not be in the future")
		assert.Equal(t, v.Valid(), false)
		assert.Equal(t, len(v.Errors), 4)
	})
}
