package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPagination(t *testing.T) {
	var c gin.Context
	req := httptest.NewRequest("GET", "/users?page=2&limit=10&orderBy=id", nil)
	c.Request = req
	res := getPagination(&c)
	if res.Page != 2 {
		t.Fatal("failed")
	}
	if res.Limit != 10 {
		t.Fatal("failed")
	}
	if res.OrderBy != "id" {
		t.Fatal("failed")
	}
}
