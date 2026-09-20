package router

import (
	"github.com/yuki-829/ToDo_GOtra1/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler.Index)
	r.GET("/ping", handler.Ping)
	r.GET("/todos", handler.ListToDos)
	r.GET("/todos/create", handler.ShowCreateForm)
	r.POST("/todos", handler.CreateTodo)
	return r
}
