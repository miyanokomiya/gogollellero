package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/responses"
)

// AuthRequired 認証フィルタ
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := handlers.GetCurrentUser(c)
		if user == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
				Key:     "invalid_auth",
				Message: "invalid auth",
			})
		}
		c.Next()
	}
}
