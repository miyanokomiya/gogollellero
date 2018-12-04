package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
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
	private := app.Group("/private")
	private.Use(AuthRequired())
	{
		private.GET("/", private1)
		private.GET("/two", private2)
	}
}

func private1(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	c.JSON(http.StatusOK, gin.H{"hello": user})
}

func private2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "auth2"})
}

// AuthRequired 認証フィルタ
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	if username == "hello" && password == "itsme" {
		session.Set("user", username) //In real world usage you'd set this to the users ID
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		log.Println(user)
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}
