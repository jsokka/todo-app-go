package db

import (
	"github.com/jsokka/todo-app-go/db/models"
)

func (db TodoAppDb) GetTags() []models.Tag {
	var tags []models.Tag
	db.Find(&tags)
	return tags
}
