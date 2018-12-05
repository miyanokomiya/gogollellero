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
	var reader *strings.Reader
	if params != nil {
		reader = strings.NewReader(params.Encode())
	}
	req := httptest.NewRequest("POST", uri, reader)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	router := server.Create()
	router.ServeHTTP(rec, req)
	return rec
}

func mockGet(uri string, params url.Values) *httptest.ResponseRecorder {
	return mockQuery("GET", uri, params)
}

func mockDelete(uri string, params url.Values) *httptest.ResponseRecorder {
	return mockQuery("DELETE", uri, params)
}

func mockQuery(method string, uri string, params url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, uri, strings.NewReader(params.Encode()))
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}
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

func TestAuthHandlerLogout(t *testing.T) {
	rec := mockDelete("/api/v1/logout", nil)

	if http.StatusOK != rec.Code {
		t.Fatal("falied")
	}
}
