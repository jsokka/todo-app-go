package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jsokka/todo-app-go/db/models"

	"gorm.io/gorm"
)

func (db TodoAppDb) baseTaskQuery() (tc *gorm.DB) {
	return db.Preload("Tags").Joins("Project")
}

func (db TodoAppDb) GetTasks() []models.Task {
	var tasks []models.Task
	db.baseTaskQuery().Find(&tasks)
	return tasks
}

func (db TodoAppDb) GetTaskById(id int) (*models.Task, error) {
	var task models.Task
	err := db.baseTaskQuery().First(&task, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrDbRecordNotFound
	}

	return &task, err
}

func (db TodoAppDb) GetTaskByIdentifier(identifier uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := db.baseTaskQuery().First(&task, "Task.Identifier=?", identifier).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrDbRecordNotFound
	}

	return &task, err
}

func (db TodoAppDb) CreateTask(task models.Task) (*models.Task, error) {
	if err := db.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (db TodoAppDb) DeleteTask(identifier string) (err error) {
	return db.Where("Identifier=?", identifier).Delete(&models.Task{}).Error
}
