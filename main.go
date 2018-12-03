package main

import (
	"github.com/miyanokomiya/gogollellero/app/server"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func main() {
	models.GormOpen()
	defer models.GormClose()
	server.Start()
}
