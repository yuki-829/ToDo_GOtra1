package router

import (
	"GO/ToDo_GOtra1/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler.Index)
	r.GET("/ping", handler.Ping)

	return r
}
