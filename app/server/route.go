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
	postsHandler := handlers.NewPostsHandler()
	tagsHandler := handlers.NewTagsHandler()

	app.Use(middlewares.CORSMiddleware())
	app.POST("/login", authHandler.Login)
	app.DELETE("/logout", authHandler.Logout)

	app.POST("/users", usersHandler.Create)
	app.GET("/tags", tagsHandler.Index)

	private := app.Group("/private")
	private.Use(middlewares.AuthRequired())
	{
		private.GET("/users", usersHandler.Index)
		private.GET("/users/:id", usersHandler.Show)
		private.PATCH("/users/:id", usersHandler.Update)
		private.DELETE("/users/:id", usersHandler.Delete)

		private.GET("/posts", postsHandler.Index)
		private.POST("/posts", postsHandler.Create)
		private.GET("/posts/:id", postsHandler.Show)
		private.PATCH("/posts/:id", postsHandler.Update)
		private.DELETE("/posts/:id", postsHandler.Delete)

		private.GET("/posts", tagsHandler.IndexOfMine)
	}
}
