package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

func TestAuthRequiredSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username", Password: "1234567890"}
		user.Create()
		defer user.Delete()

		login(h.eng, user.ID)
		h.eng.Use(server.AuthRequired())
		h.eng.GET("/users/:name", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", user.Name), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestAuthRequiredFailed(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.Use(server.AuthRequired())
		h.eng.GET("/users/:name", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		req := httptest.NewRequest("GET", "/users/name", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestAuthRequiredFailedWhenUserNotFound(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username", Password: "1234567890"}

		login(h.eng, user.ID)
		h.eng.Use(server.AuthRequired())
		h.eng.GET("/users/:name", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		req := httptest.NewRequest("GET", fmt.Sprintf("/users/%s", user.Name), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}
