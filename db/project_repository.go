package db

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/jsokka/todo-app-go/db/models"
	"gorm.io/gorm"
)

func (db TodoAppDb) GetProjects() []models.Project {
	var projects []models.Project
	db.Find(&projects)
	return projects
}

func (db TodoAppDb) GetProjectWithIdentifier(id uuid.UUID) (*models.Project, error) {
	var project models.Project
	err := db.First(&project, "Identifier=?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrDbRecordNotFound
	}

	return &project, err
}

func (db TodoAppDb) GetProjectIdByIdentifier(identifier string) (sql.NullInt32, error) {
	var project models.Project

	if strings.TrimSpace(identifier) == "" {
		return sql.NullInt32{}, nil
	}

	if err := db.Select("Id").First(&project, "Identifier=?", identifier).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return sql.NullInt32{}, ErrDbRecordNotFound
	}

	return sql.NullInt32{Int32: int32(project.Id), Valid: true}, nil
}

func (db TodoAppDb) CreateProject(project models.Project) (*models.Project, error) {
	if err := db.Create(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (db TodoAppDb) DeleteProject(identifier string) (err error) {
	return db.Where("Identifier=?", identifier).Delete(&models.Project{}).Error
}
