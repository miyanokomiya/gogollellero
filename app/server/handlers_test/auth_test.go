package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func mockPost(uri string, params url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", uri, strings.NewReader(params.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := server.Create()
	router.ServeHTTP(rec, req)
	return rec
}

func TestAuthHandlerLogin1(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	user.SetPassword("password")
	user.Create()
	defer user.Delete()

	values := url.Values{}
	values.Add("username", "username")
	values.Add("password", "password")
	rec := mockPost("/api/v1/login", values)

	if http.StatusOK != rec.Code {
		t.Fatal("falied")
	}
}

func TestAuthHandlerLogin2(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	user.SetPassword("password")
	user.Create()
	defer user.Delete()

	values := url.Values{}
	values.Add("username", "username")
	values.Add("password", "invalid")
	rec := mockPost("/api/v1/login", values)

	if http.StatusUnauthorized != rec.Code {
		t.Fatal("falied")
	}
}

func TestAuthHandlerLogin3(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	user.SetPassword("password")
	user.Create()
	defer user.Delete()

	values := url.Values{}
	values.Add("username", "unknown")
	values.Add("password", "password")
	rec := mockPost("/api/v1/login", values)

	if http.StatusUnauthorized != rec.Code {
		t.Fatal("falied")
	}
}
