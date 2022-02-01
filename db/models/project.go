package models

import (
	"database/sql"
)

type Project struct {
	DbModel
	Name        string
	Description sql.NullString
	Deadline    sql.NullTime
}
