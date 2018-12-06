package handlers

import (
	"fmt"
	"net/http"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

// UsersHandler ユーザーハンドラのインタフェース
type UsersHandler interface {
	Show(c *gin.Context)
	Create(c *gin.Context)
}

// NewUsersHandler 生成
func NewUsersHandler() UsersHandler {
	return &usersHandler{}
}

type usersHandler struct{}

// Show 詳細
func (h *usersHandler) Show(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s", name),
	})
}

// Create 作成
func (h *usersHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := models.User{Name: name}
	err := user.SetPassword(password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "empty_password",
			Message: "empty password",
		})
		return
	}
	err = user.Create()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "validation_error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &user)
}
