package handlers_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func TestUsersHandlerCreateSuccess(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	defer func() {
		user.Read()
		user.Delete()
	}()

	json := handlers.CraeteJSON{
		Name:     user.Name,
		Password: "password",
	}
	rec := mockPost("/api/v1/users", json)

	if http.StatusOK != rec.Code {
		t.Fatal("falied", rec)
	}
}

func TestUsersHandlerCreateFailed(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	defer func() {
		user.Read()
		user.Delete()
	}()

	values := url.Values{}
	values.Add("password", "password")

	json := handlers.CraeteJSON{
		Password: "password",
	}
	rec := mockPost("/api/v1/users", json)

	if http.StatusOK == rec.Code {
		t.Fatal("falied", rec)
	}
}
