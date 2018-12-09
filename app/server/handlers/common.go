package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/constants"
	"github.com/miyanokomiya/gogollellero/app/server/models"
	"github.com/miyanokomiya/gogollellero/app/server/responses"
)

func getPagination(c *gin.Context) *models.Pagination {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 50
	}
	orderBy := c.Query("orderBy")
	return &models.Pagination{
		Page:    page,
		Limit:   limit,
		OrderBy: orderBy,
	}
}

// GetCurrentUser ログインユーザー取得
func GetCurrentUser(c *gin.Context) *models.User {
	session := sessions.Default(c)
	v := session.Get(constants.SessionUser)
	if v != nil {
		if id, ok := v.(int); ok {
			user := models.User{}
			user.ID = int(id)
			err := user.Read()
			if err == nil {
				return &user
			}
			log.Println(err)
		}
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, responses.Error{
		Key:     "invalid_auth",
		Message: "invalid auth",
	})
	return nil
}
