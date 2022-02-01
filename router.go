package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jsokka/todo-app-go/db"
	handlers "github.com/jsokka/todo-app-go/handlers"
)

func InitRouter() *gin.Engine {
	h := handlers.New(&db.Db)

	router := gin.Default()
	h.InitProjectsHandler(router)
	h.InitTasksHandler(router)
	h.InitTagsHandler(router)

	return router
}
