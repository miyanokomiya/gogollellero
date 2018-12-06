package handlerstest

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func TestUsersHandlerCreate1(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	defer func() {
		user.Read()
		user.Delete()
	}()

	values := url.Values{}
	values.Add("name", user.Name)
	values.Add("password", "password")
	rec := mockPost("/api/v1/users", values)

	if http.StatusOK != rec.Code {
		t.Fatal("falied", rec)
	}
}

func TestUsersHandlerCreate2(t *testing.T) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	defer func() {
		user.Read()
		user.Delete()
	}()

	values := url.Values{}
	values.Add("password", "password")
	rec := mockPost("/api/v1/users", values)

	if http.StatusOK == rec.Code {
		t.Fatal("falied", rec)
	}
}
