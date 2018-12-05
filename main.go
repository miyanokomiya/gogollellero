package main

import (
	"os"

	"github.com/miyanokomiya/gogollellero/app/server"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func main() {
	models.GormOpen()
	defer models.GormClose()
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	server.Create().Run(":" + port)
}
