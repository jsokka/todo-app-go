package models

import (
	"time"

	dbmodels "github.com/jsokka/todo-app-go/db/models"
)

type Task struct {
	Identifier  string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Priority    dbmodels.TaskPriority `json:"priority"`
	Deadline    time.Time             `json:"deadline"`
	CompletedOn time.Time             `json:"completed_on"`
	ProjectId   string                `json:"project_id"`
	Tags        []Tag                 `json:"tags"`
}

func NewTask(identifier string, projectId string, title string, description string, priority dbmodels.TaskPriority,
	deadline time.Time, completedOn time.Time, tags []Tag) Task {
	return Task{
		Identifier:  identifier,
		Title:       title,
		Description: description,
		Priority:    priority,
		Deadline:    deadline,
		CompletedOn: completedOn,
		ProjectId:   projectId,
		Tags:        tags,
	}
}
