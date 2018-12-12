package handlers

import (
	"net/http"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

// TagsHandler タグハンドラのインタフェース
type TagsHandler interface {
	Index(c *gin.Context)
	IndexOfMine(c *gin.Context)
}

// NewTagsHandler 生成
func NewTagsHandler() TagsHandler {
	return &tagsHandler{}
}

type tagsHandler struct{}

// Index 一覧
func (h *tagsHandler) Index(c *gin.Context) {
	pagenation := getPagination(c)
	tags := models.Tags{}
	if err := tags.Index(pagenation); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Key:     "internal_server_error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, tags)
}

// IndexOfMine 利用中タグ一覧
func (h *tagsHandler) IndexOfMine(c *gin.Context) {
	user := GetCurrentUser(c)
	if user == nil {
		return
	}
	tags := models.Tags{}
	if err := tags.IndexOfUser(user.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Key:     "internal_server_error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, tags)
}
