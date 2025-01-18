package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Patients PatientModel
	Users    UserModel
	Tokens   TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Patients: PatientModel{DB: db},
		Users:    UserModel{DB: db},
		Tokens:   TokenModel{DB: db},
	}
}
