package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsokka/todo-app-go/db"
	"github.com/jsokka/todo-app-go/models"
)

type handler struct {
	DB *db.TodoAppDb
}

func New(db *db.TodoAppDb) handler {
	return handler{db}
}

func (h handler) HandleErrorAndAbort(c *gin.Context, err error) {
	switch err {
	case db.ErrDbRecordNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, models.NewError(err.Error(), http.StatusNotFound, ""))
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			models.NewError("Internal server error", http.StatusInternalServerError, err.Error()))
	}
}

func (h handler) CreatedAtJSON(c *gin.Context, model interface{}, location string) {
	c.Writer.Header().Set("Content-Location", location)
	c.JSON(http.StatusCreated, model)
}

func (h handler) AbortBadRequest(c *gin.Context, details string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, models.NewError("Invalid data", http.StatusBadRequest, details))
}
