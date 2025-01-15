package data

import "database/sql"

type Models struct {
	Patients PatientModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Patients: PatientModel{DB: db},
	}
}
