package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	dbmodels "github.com/jsokka/todo-app-go/db/models"
	"github.com/jsokka/todo-app-go/models"
)

func (h handler) InitProjectsHandler(router *gin.Engine) {
	router.GET("/projects", h.GetProjects)
	router.GET("/projects/:id", h.GetProject)
	router.POST("/projects", h.PostProject)
	router.DELETE("/projects/:id", h.DeleteProject)
}

func (h handler) GetProjects(c *gin.Context) {
	c.JSON(http.StatusOK, projectsFromDbModels(h.DB.GetProjects()))
}

func (h handler) GetProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		h.AbortBadRequest(c, "Id is not a valid UUID")
		return
	}

	project, err := h.DB.GetProjectWithIdentifier(id)

	if err != nil {
		h.HandleErrorAndAbort(c, err)
		return
	}

	c.JSON(http.StatusOK, projectFromDbModel(*project))
}

func (h handler) PostProject(c *gin.Context) {
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		h.AbortBadRequest(c, "")
		return
	}

	newProject, err := h.DB.CreateProject(projectToDbModel(project))

	if err != nil {
		h.HandleErrorAndAbort(c, err)
		return
	}

	project = projectFromDbModel(*newProject)

	h.CreatedAtJSON(c, project, fmt.Sprintf("/projects/%s", project.Identifier))
}

func (h handler) DeleteProject(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		h.AbortBadRequest(c, "Id is not a valid UUID")
		return
	}

	if err := h.DB.DeleteProject(id.String()); err != nil {
		h.HandleErrorAndAbort(c, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func projectsFromDbModels(sourceProjects []dbmodels.Project) []models.Project {
	var ret []models.Project
	for _, p := range sourceProjects {
		ret = append(ret, projectFromDbModel(p))
	}
	return ret
}

func projectFromDbModel(source dbmodels.Project) models.Project {
	return models.NewProject(
		source.Identifier,
		source.Name,
		source.Description.String,
		source.Deadline.Time)
}

func projectToDbModel(source models.Project) dbmodels.Project {
	description := sql.NullString{String: source.Description, Valid: len(source.Description) > 0}
	deadline := sql.NullTime{Time: source.Deadline, Valid: !source.Deadline.IsZero()}
	return dbmodels.Project{
		Name:        source.Name,
		Description: description,
		Deadline:    deadline,
	}
}
