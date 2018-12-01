package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/handlers"
)

// RouteV1 API v1 ルーティング
func RouteV1(app *gin.Engine) {
	helloHandler := handlers.NewHelloHandler()
	apiGroup := app.Group("api/v1")
	{
		apiGroup.GET("/user/:name", helloHandler.Show)
	}
}
