package handlers

import (
	"log"
	"net/http"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/constants"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

// NewUsersHandler 生成
func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

type authHandler struct{}

// Login ログイン
func (h *authHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{Name: username}
	err := user.Read()
	log.Println(user)
	if err != nil {
		log.Println(err)
		respondFailedLogin(c)
		return
	}
	if !user.Authenticate(password) {
		log.Println("Invalid password")
		respondFailedLogin(c)
		return
	}

	session := sessions.Default(c)
	session.Set(constants.SessionUser, user.ID)
	err = session.Save()
	if err != nil {
		log.Println(err)
		respondFailedLogin(c)
		return
	}

	log.Println("Success login", user.ID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

// Logout ログアウト
func (h *authHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(constants.SessionUser)
	if userID == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
		return
	}
	session.Delete(constants.SessionUser)
	err := session.Save()
	if err != nil {
		log.Println("Falied delete session")
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Error{
			Key:     "failed_logout",
			Message: "failed logout",
		})
		return
	}
	log.Println("Success logout", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func respondFailedLogin(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Error{
		Key:     "failed_login",
		Message: "failed login",
	})
}
