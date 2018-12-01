package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandlerInterface インタフェース
type HelloHandlerInterface interface {
	GetName(c *gin.Context)
}

// NewHelloHandler 生成
func NewHelloHandler() HelloHandlerInterface {
	return &helloHandler{}
}

type helloHandler struct {
}

func (h *helloHandler) GetName(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello %s", name),
	})
}
