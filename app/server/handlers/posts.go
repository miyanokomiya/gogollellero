package handlers

import (
	"net/http"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

// PostsHandler ユーザーハンドラのインタフェース
type PostsHandler interface {
	Index(c *gin.Context)
}

// NewPostsHandler 生成
func NewPostsHandler() PostsHandler {
	return &postsHandler{}
}

type postsHandler struct{}

// Index 一覧
func (h *postsHandler) Index(c *gin.Context) {
	user := GetCurrentUser(c)
	if user == nil {
		return
	}
	pagenation := getPagination(c)
	posts := models.Posts{}
	if err := posts.IndexInUser(pagenation, user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Key:     "internal_server_error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}
