package server

import (
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/middlewares"
)

// RouteV1 API v1 ルーティング
func RouteV1(app *gin.Engine) {
	authHandler := handlers.NewAuthHandler()
	usersHandler := handlers.NewUsersHandler()

	app.Use(middlewares.CORSMiddleware())
	app.POST("/login", authHandler.Login)
	app.DELETE("/logout", authHandler.Logout)

	app.POST("/users", usersHandler.Create)

	private := app.Group("/private")
	private.Use(middlewares.AuthRequired())
	{
		private.GET("/users", usersHandler.Index)
		private.GET("/users/:id", usersHandler.Show)
		private.PATCH("/users/:id", usersHandler.Update)
		private.DELETE("/users/:id", usersHandler.Delete)
	}
}
