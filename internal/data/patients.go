package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/validator"
)

type Patient struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Notes       string    `json:"notes"`
}

func ValidatePatient(v *validator.Validator, patient *Patient) {
	v.Check(patient.FirstName != "", "first_name", "must be provided")
	v.Check(len(patient.FirstName) <= 50, "first_name", "must not be more than 50 bytes long")

	v.Check(patient.LastName != "", "last_name", "must be provided")
	v.Check(len(patient.LastName) <= 50, "last_name", "must not be more than 50 bytes long")

	v.Check(patient.Gender != "", "gender", "must be provided")
	v.Check(validator.PermittedValue(patient.Gender, "M", "F"), "gender", "must be M or F only")

	v.Check(patient.DateOfBirth.After(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)), "date_of_birth", "must not be before 1900-01-01")
	v.Check(patient.DateOfBirth.Before(time.Now()), "date_of_birth", "must not be in the future")

}

type PatientModel struct {
	DB *sql.DB
}

func (m PatientModel) Insert(patient *Patient) error {
	query := `
	INSERT INTO patients (first_name, last_name, date_of_birth, gender)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at`

	args := []any{patient.FirstName, patient.LastName, patient.DateOfBirth, patient.Gender}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&patient.ID, &patient.CreatedAt)
}
