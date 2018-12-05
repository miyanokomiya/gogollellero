package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersHandler ユーザーハンドラのインタフェース
type UsersHandler interface {
	Show(c *gin.Context)
}

// NewUsersHandler 生成
func NewUsersHandler() UsersHandler {
	return &usersHandler{}
}

type usersHandler struct{}

func (h *usersHandler) Show(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s", name),
	})
}
