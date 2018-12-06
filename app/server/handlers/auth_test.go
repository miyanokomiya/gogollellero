package handlers_test

import (
	"net/http"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func TestAuthHandlerLogin1(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	user.SetPassword("password")
	user.Create()
	defer user.Delete()

	json := handlers.LoginJson{
		UserName: user.Name,
		Password: "password",
	}
	rec := mockPost("/api/v1/login", json)

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

	json := handlers.LoginJson{
		UserName: user.Name,
		Password: "invalid",
	}
	rec := mockPost("/api/v1/login", json)

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

	json := handlers.LoginJson{
		UserName: "unknown",
		Password: "password",
	}
	rec := mockPost("/api/v1/login", json)

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
