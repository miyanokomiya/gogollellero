package handlers_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func TestUsersHandlerShowSuccess(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username", Password: "1234567890"}
	user.Create()
	defer user.Delete()

	rec := mockGet(fmt.Sprintf("/api/v1/users/%d", user.ID), nil)

	if http.StatusOK != rec.Code {
		t.Fatal("falied", rec)
	}
}

func TestUsersHandlerShowFailed(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username", Password: "1234567890"}
	user.Create()
	defer user.Delete()

	rec := mockGet(fmt.Sprintf("/api/v1/users/%d", user.ID+1), nil)

	if http.StatusOK == rec.Code {
		t.Fatal("falied", rec)
	}
}

func TestUsersHandlerCreateSuccess(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	defer func() {
		user.Read()
		user.Delete()
	}()

	json := handlers.UserCraeteJSON{
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

	json := handlers.UserCraeteJSON{
		Password: "password",
	}
	rec := mockPost("/api/v1/users", json)

	if http.StatusOK == rec.Code {
		t.Fatal("falied", rec)
	}
}

func TestUsersHandlerDeleteSuccess(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username", Password: "1234567890"}
	user.Create()
	defer user.Delete()

	rec := mockDelete(fmt.Sprintf("/api/v1/users/%d", user.ID), nil)

	if http.StatusOK != rec.Code {
		t.Fatal("falied", rec)
	}
}

func TestUsersHandlerDeleteFailed(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username", Password: "1234567890"}
	user.Create()
	defer user.Delete()

	rec := mockDelete(fmt.Sprintf("/api/v1/users/%d", user.ID+1), nil)

	if http.StatusOK == rec.Code {
		t.Fatal("falied", rec)
	}
}
