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
	Create(c *gin.Context)
}

// NewPostsHandler 生成
func NewPostsHandler() PostsHandler {
	return &postsHandler{}
}

type postsHandler struct{}

// Index 一覧 ログイン者に属するもの
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

// PostCraeteJSON Createパラメータ
type PostCraeteJSON struct {
	Title    string `json:"title" binding:"required,lte=256"`
	Problem  string `json:"problem"`
	Solution string `json:"solution"`
	Lesson   string `json:"lesson"`
}

// Create 作成
func (h *postsHandler) Create(c *gin.Context) {
	user := GetCurrentUser(c)
	if user == nil {
		return
	}
	json := PostCraeteJSON{}
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}

	post := models.Post{
		UserID:   user.ID,
		User:     *user,
		Title:    json.Title,
		Problem:  json.Problem,
		Solution: json.Solution,
		Lesson:   json.Lesson,
	}

	if err := post.Create(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "validation_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &post)
}
