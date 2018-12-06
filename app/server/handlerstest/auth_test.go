package handlerstest

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/models"
)

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
