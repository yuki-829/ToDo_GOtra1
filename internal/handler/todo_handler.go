package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuki-829/ToDo_GOtra1/internal/service"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"Title": "ToDo App",
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ListToDos(c *gin.Context) {
	todos, _ := service.GetTodos()
	c.HTML(200, "todo/list.tmpl", gin.H{
		"Todos": todos,
	})
}

func CreateTodo(c *gin.Context) {
	task := c.PostForm("task")
	service.CreateTodo(task)
	c.Redirect(302, "/todos")
}

func ShowCreateForm(c *gin.Context) {
	c.HTML(200, "todo/create.tmpl", nil)
}
