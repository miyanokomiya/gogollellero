package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/constants"
)

// Login ログイン
func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	if username == "hello" && password == "itsme" {
		session.Set(constants.SessionUser, username) //In real world usage you'd set this to the users ID
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

// Logout ログアウト
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(constants.SessionUser)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		log.Println(user)
		session.Delete(constants.SessionUser)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}
