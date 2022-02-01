package models

import (
	"database/sql"
)

type Task struct {
	DbModel
	Title       string
	Description sql.NullString
	Priority    TaskPriority
	Deadline    sql.NullTime
	CompletedOn sql.NullTime
	ProjectId   sql.NullInt32
	Tags        []Tag   `gorm:"many2many:TaskTag;foreignKey:Id;joinForeignKey:TaskId;References:Id;joinReferences:TagId"`
	Project     Project `gorm:"foreignKey:ProjectId"`
}

type TaskPriority string

const (
	Normal   TaskPriority = "Normal"
	Low      TaskPriority = "Low"
	High     TaskPriority = "High"
	VeryHigh TaskPriority = "VeryHigh"
)
