package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

var authHandlers = handlers.NewAuthHandler()

func TestAuthHandlerLogin1(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username"}
		user.SetPassword("password")
		user.Create()
		defer user.Delete()
		h.eng.POST("/login", authHandlers.Login)
		req := httptest.NewRequest("POST", "/login", createJsonParams(handlers.LoginJson{
			UserName: user.Name,
			Password: "password",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestAuthHandlerLogin2(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username"}
		user.SetPassword("password")
		user.Create()
		defer user.Delete()
		h.eng.POST("/login", authHandlers.Login)
		req := httptest.NewRequest("POST", "/login", createJsonParams(handlers.LoginJson{
			UserName: user.Name,
			Password: "invalid",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestAuthHandlerLogin3(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username"}
		user.SetPassword("password")
		user.Create()
		defer user.Delete()
		h.eng.POST("/login", authHandlers.Login)
		req := httptest.NewRequest("POST", "/login", createJsonParams(handlers.LoginJson{
			UserName: "unkown",
			Password: "password",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestAuthHandlerLogout(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.DELETE("/logout", authHandlers.Logout)
		req := httptest.NewRequest("DELETE", "/logout", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}
