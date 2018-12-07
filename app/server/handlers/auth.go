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

// AuthHandler 認証ハンドラ
type AuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

// NewAuthHandler 生成
func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

type authHandler struct{}

// LoginJSON ログインJSON
type LoginJSON struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login ログイン
func (h *authHandler) Login(c *gin.Context) {
	json := LoginJSON{}
	if err := c.BindJSON(&json); err != nil {
		log.Println(err)
		respondFailedLogin(c)
		return
	}

	user := models.User{Name: json.UserName}
	if err := user.Read(); err != nil {
		log.Println(err)
		respondFailedLogin(c)
		return
	}

	if !user.Authenticate(json.Password) {
		log.Println("Invalid password")
		respondFailedLogin(c)
		return
	}

	session := sessions.Default(c)
	session.Set(constants.SessionUser, user.ID)
	if err := session.Save(); err != nil {
		log.Println(err)
		respondFailedLogin(c)
		return
	}

	log.Println("Success login", user.ID)
	c.JSON(http.StatusOK, user)
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
