package models

import "time"

type Project struct {
	Identifier  string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

func NewProject(identifier string, name string, description string, deadline time.Time) Project {
	return Project{Identifier: identifier, Name: name, Description: description, Deadline: deadline}
}
