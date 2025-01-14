package data

import "time"

type Patient struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Gender      string    `json:"gender"`
	Notes       string    `json:"notes"`
}

func ValidatePatient(patient *Patient) map[string]string {
	errors := make(map[string]string)
	if patient.FirstName == "" {
		errors["first_name"] = "must be provided"
	}

	if patient.LastName == "" {
		errors["last_name"] = "must be provided"
	}

	return errors
}
