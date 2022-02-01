package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	dbmodels "github.com/jsokka/todo-app-go/db/models"
	models "github.com/jsokka/todo-app-go/models"
)

func (h handler) InitTasksHandler(router *gin.Engine) {
	router.GET("/tasks", h.GetTasks)
	router.GET("/tasks/:id", h.GetTask)
	router.POST("/tasks", h.PostTask)
	router.DELETE("/tasks/:id", h.DeleteTask)
}

func (h handler) GetTasks(c *gin.Context) {
	tasks := h.DB.GetTasks()
	c.JSON(http.StatusOK, tasksFromDbModels(tasks))
}

func (h handler) GetTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		h.AbortBadRequest(c, "Id is not a valid UUID")
		return
	}

	task, err := h.DB.GetTaskByIdentifier(id)

	if err != nil {
		h.HandleErrorAndAbort(c, err)
		return
	}

	c.JSON(http.StatusOK, taskFromDbModel(*task))
}

func (h handler) PostTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		h.AbortBadRequest(c, "")
		return
	}

	projectId, err := h.DB.GetProjectIdByIdentifier(task.ProjectId)

	if err != nil {
		h.AbortBadRequest(c, err.Error())
		return
	}

	newTask, err := h.DB.CreateTask(taskToDbModel(task, projectId))

	if err != nil {
		h.HandleErrorAndAbort(c, err)
		return
	}

	newTask, _ = h.DB.GetTaskById(newTask.Id)
	task = taskFromDbModel(*newTask)

	h.CreatedAtJSON(c, task, fmt.Sprintf("/tasks/%s", task.Identifier))
}

func (h handler) DeleteTask(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		h.AbortBadRequest(c, "Id is not a valid UUID")
		return
	}

	if err := h.DB.DeleteTask(id.String()); err != nil {
		h.HandleErrorAndAbort(c, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func tasksFromDbModels(sourceTasks []dbmodels.Task) []models.Task {
	var ret []models.Task
	for _, t := range sourceTasks {
		ret = append(ret, taskFromDbModel(t))
	}
	return ret
}

func taskFromDbModel(source dbmodels.Task) models.Task {
	return models.NewTask(
		source.Identifier,
		source.Project.Identifier,
		source.Title,
		source.Description.String,
		source.Priority,
		source.Deadline.Time,
		source.CompletedOn.Time,
		tagsFromDbModels(source.Tags),
	)
}

func taskToDbModel(source models.Task, projectId sql.NullInt32) dbmodels.Task {
	description := sql.NullString{String: source.Description, Valid: len(source.Description) > 0}
	deadline := sql.NullTime{Time: source.Deadline, Valid: !source.Deadline.IsZero()}
	completedOn := sql.NullTime{Time: source.CompletedOn, Valid: !source.CompletedOn.IsZero()}
	return dbmodels.Task{
		Title:       source.Title,
		Description: description,
		Priority:    source.Priority,
		Deadline:    deadline,
		CompletedOn: completedOn,
		ProjectId:   projectId,
	}
}
