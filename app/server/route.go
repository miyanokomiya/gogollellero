package server

import (
	"net/http"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/miyanokomiya/gogollellero/app/server/constants"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/handlers"
)

// RouteV1 API v1 ルーティング
func RouteV1(app *gin.Engine) {
	usersHandler := handlers.NewUsersHandler()
	apiGroup := app.Group("api/v1")
	{
		apiGroup.GET("/users/:name", usersHandler.Show)
	}
	private := app.Group("/private")
	private.Use(AuthRequired())
	{
		private.GET("/users/:name", private2)
	}
}

func private2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "auth2"})
}

// AuthRequired 認証フィルタ
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(constants.SessionUser)
		if user == nil {
			// You'd normally redirect to login page
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
				Key:     "invalid_auth",
				Message: "invalud auth",
			})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}
