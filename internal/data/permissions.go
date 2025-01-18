package data

import (
	"database/sql"
	"slices"
)

type permissions []string

func (p permissions) Include(code string) bool {
	return slices.Contains(p, code)
}

type PermissionModel struct {
	DB *sql.DB
}

func (m PermissionModel) 
