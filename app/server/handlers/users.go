package handlers

import (
	"log"
	"net/http"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

// UsersHandler ユーザーハンドラのインタフェース
type UsersHandler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// NewUsersHandler 生成
func NewUsersHandler() UsersHandler {
	return &usersHandler{}
}

type usersHandler struct{}

// Index 一覧
func (h *usersHandler) Index(c *gin.Context) {
	pagenation := getPagination(c)
	log.Println("-----------------", pagenation)
	users := models.Users{}
	if err := users.Index(pagenation); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Key:     "internal_server_error",
			Message: "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Show 詳細
func (h *usersHandler) Show(c *gin.Context) {
	id := parseID(c)
	if id == 0 {
		return
	}

	user := getUser(c, id)
	if user == nil {
		return
	}

	c.JSON(http.StatusOK, user)
}

// UserCraeteJSON Createパラメータ
type UserCraeteJSON struct {
	Name     string `json:"name" binding:"required,gte=4,lte=64"`
	Password string `json:"password" binding:"required,gte=8,lte=64"`
}

// Create 作成
func (h *usersHandler) Create(c *gin.Context) {
	json := UserCraeteJSON{}
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}

	user := models.User{Name: json.Name}
	if err := user.SetPassword(json.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "empty_password",
			Message: "empty password",
		})
		return
	}

	if err := user.Create(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "validation_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &user)
}

// UserUpdateJSON Updateパラメータ
type UserUpdateJSON struct {
	Name string `json:"name" binding:"gte=4,lte=64"`
}

// Update 作成
func (h *usersHandler) Update(c *gin.Context) {
	id := parseID(c)
	if id == 0 {
		return
	}

	json := UserUpdateJSON{}
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}

	user := getUser(c, id)
	if user == nil {
		return
	}

	user.Name = json.Name
	if err := user.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "validation_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &user)
}

// Delete 削除
func (h *usersHandler) Delete(c *gin.Context) {
	id := parseID(c)
	if id == 0 {
		return
	}

	user := getUser(c, id)
	if user == nil {
		return
	}

	if err := user.Delete(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Key:     "failed_delete_user",
			Message: "failed delete user",
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func getUser(c *gin.Context, id int) *models.User {
	user := models.User{}
	user.ID = id
	if err := user.Read(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Error{
			Key:     "not_found_user",
			Message: "not found user",
		})
		return nil
	}
	return &user
}
