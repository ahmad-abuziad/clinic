package data

import (
	"time"

	"github.com/ahmad-abuziad/clinic/internal/validator"
)

type Patient struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Notes       string    `json:"notes"`
}

func ValidatePatient(v *validator.Validator, patient *Patient) {
	v.Check(patient.FirstName != "", "first_name", "must be provided")
	v.Check(len(patient.FirstName) <= 50, "first_name", "must not be more than 50 bytes long")

	v.Check(patient.LastName != "", "last_name", "must be provided")
	v.Check(len(patient.LastName) <= 50, "last_name", "must not be more than 50 bytes long")
}
