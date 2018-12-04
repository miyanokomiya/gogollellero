package main

import (
	"github.com/miyanokomiya/gogollellero/app/server"
)

func main() {
	// models.GormOpen()
	// defer models.GormClose()
	server.Start()
}
