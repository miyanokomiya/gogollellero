package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/models"
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
