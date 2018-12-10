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
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

// Show 詳細
func (h *postsHandler) Show(c *gin.Context) {
	id := parseID(c)
	if id == 0 {
		return
	}

	post := getPost(c, id)
	if post == nil {
		return
	}

	c.JSON(http.StatusOK, post)
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

// PostUpdateJSON Updateパラメータ
type PostUpdateJSON struct {
	Title    *string  `json:"title" binding:"required,lte=256"`
	Problem  *string  `json:"problem"`
	Solution *string  `json:"solution"`
	Lesson   *string  `json:"lesson"`
	Tags     []string `json:"tags"`
}

// Update 作成
func (h *postsHandler) Update(c *gin.Context) {
	id := parseID(c)
	if id == 0 {
		return
	}

	json := PostUpdateJSON{}
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}

	post := getPost(c, id)
	if post == nil {
		return
	}

	if json.Title != nil {
		post.Title = *json.Title
	}
	if json.Problem != nil {
		post.Problem = *json.Problem
	}
	if json.Solution != nil {
		post.Solution = *json.Solution
	}
	if json.Lesson != nil {
		post.Lesson = *json.Lesson
	}
	if json.Tags != nil {
		tags, err := models.CreateTagsIfNotExist(json.Tags)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
				Key:     "failed_update_post",
				Message: "failed update post",
			})
			return
		}
		post.Tags = tags
	}
	if err := post.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "validation_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &post)
}

// Delete 削除
func (h *postsHandler) Delete(c *gin.Context) {
	id := parseID(c)
	if id == 0 {
		return
	}

	post := getPost(c, id)
	if post == nil {
		return
	}

	if err := post.Delete(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Key:     "failed_delete_post",
			Message: "failed delete post",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func getPost(c *gin.Context, id int) *models.Post {
	post := models.Post{}
	post.ID = id
	if err := post.Read(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Error{
			Key:     "not_found_post",
			Message: "not found post",
		})
		return nil
	}
	return &post
}
