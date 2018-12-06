package handlers

import (
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

// ShowParams Showパラメータ
type ShowParams struct {
	Name string `json:"name" binding:"required,gte=4,lte=64"`
}

// Show 詳細
func (h *usersHandler) Show(c *gin.Context) {
	params := ShowParams{}
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
			Key:     "invalid_params",
			Message: "invalid params",
		})
		return
	}
	user := models.User{Name: params.Name}
	if err := user.Read(); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, responses.Error{
			Key:     "not_found_user",
			Message: "not found user",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CraeteJSON Createパラメータ
type CraeteJSON struct {
	Name     string `json:"name" binding:"required,gte=4,lte=64"`
	Password string `json:"password" binding:"required,gte=8,lte=64"`
}

// Create 作成
func (h *usersHandler) Create(c *gin.Context) {
	json := CraeteJSON{}
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
