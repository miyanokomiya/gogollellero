package server

import (
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Start 起動
func Start() {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	RouteV1(r)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(":" + port)
}
