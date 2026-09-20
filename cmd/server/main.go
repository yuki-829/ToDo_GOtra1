package main

import (
	"GO/ToDo_GOtra1/internal/router"
)

func main() {
	r := router.InitRouter()
	r.Run()
}
