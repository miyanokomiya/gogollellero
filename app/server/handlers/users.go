package handlers

import (
	"net/http"
	"strconv"

	"github.com/miyanokomiya/gogollellero/app/server/responses"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

// UsersHandler ユーザーハンドラのインタフェース
type UsersHandler interface {
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

// Show 詳細
func (h *usersHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}
	user := models.User{}
	user.ID = id
	if err := user.Read(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Error{
			Key:     "not_found_user",
			Message: "not found user",
		})
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
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"gte=4,lte=64"`
}

// Update 作成
func (h *usersHandler) Update(c *gin.Context) {
	json := UserUpdateJSON{}
	if err := c.BindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}

	user := models.User{}
	user.ID = json.ID
	if err := user.Read(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Error{
			Key:     "not_found_user",
			Message: "not found user",
		})
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}
	user := models.User{}
	user.ID = id
	if err := user.Read(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Error{
			Key:     "not_found_user",
			Message: "not found user",
		})
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
