package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dbmodels "github.com/jsokka/todo-app-go/db/models"
	models "github.com/jsokka/todo-app-go/models"
)

func (h handler) InitTagsHandler(router *gin.Engine) {
	router.GET("tags", h.GetTags)
}

func (h handler) GetTags(c *gin.Context) {
	tags := h.DB.GetTags()
	c.JSON(http.StatusOK, tagsFromDbModels(tags))
}

func tagsFromDbModels(sourceTags []dbmodels.Tag) []models.Tag {
	ret := make([]models.Tag, 0)
	for _, t := range sourceTags {
		ret = append(ret, tagFromDbModel(t))
	}
	return ret
}

func tagFromDbModel(source dbmodels.Tag) models.Tag {
	return models.Tag{Identifier: source.Identifier, Name: source.Name, CreatedOn: source.CreatedOn}
}

func tagToDbModel(source models.Tag) dbmodels.Tag {
	return dbmodels.Tag{Name: source.Name}
}
