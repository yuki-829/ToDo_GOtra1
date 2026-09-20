package main

import (
	"github.com/yuki-829/ToDo_GOtra1/internal/db"
	"github.com/yuki-829/ToDo_GOtra1/internal/model"
	"github.com/yuki-829/ToDo_GOtra1/internal/router"
)

func main() {
	db.Init()
	db.DB.AutoMigrate(&model.Todo{})
	r := router.InitRouter()
	r.LoadHTMLGlob("../../web/templates/**/*.tmpl")
	r.Run()
}
