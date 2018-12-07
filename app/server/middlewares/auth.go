package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/constants"
	"github.com/miyanokomiya/gogollellero/app/server/models"
	"github.com/miyanokomiya/gogollellero/app/server/responses"
)

// AuthRequired 認証フィルタ
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get(constants.SessionUser)
		if v != nil {
			if id, ok := v.(int); ok {
				user := models.User{}
				user.ID = int(id)
				err := user.Read()
				if err == nil {
					c.Next()
					return
				}
				log.Println(err)
			}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_auth",
			Message: "invalud auth",
		})
	}
}
