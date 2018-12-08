package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miyanokomiya/gogollellero/app/server/handlers"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

var usersHandlers = handlers.NewUsersHandler()

func TestUsersHandlerIndexSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		userListWrapper(10, func(users models.Users) {
			h.eng.GET("/users", usersHandlers.Index)
			req := httptest.NewRequest("GET", "/users?page=2&limit=3&orderBy=name", nil)
			h.eng.ServeHTTP(h.rec, req)

			if http.StatusOK != h.rec.Code {
				t.Fatal("falied", h.rec)
			}
			var resUsers models.Users
			if err := json.Unmarshal(h.rec.Body.Bytes(), &resUsers); err != nil {
				t.Fatal("falied", h.rec)
			}
			if len(resUsers) != 3 {
				t.Fatal("falied", h.rec)
			}
			if resUsers[0].Name != "user_3" {
				t.Fatal("falied", h.rec)
			}
		})
	})
}

func TestUsersHandlerShowSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username", Password: "1234567890"}
		user.Create()
		defer user.Delete()
		h.eng.GET("/users/:id", usersHandlers.Show)
		req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerShowFailed(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username", Password: "1234567890"}
		user.Create()
		defer user.Delete()
		h.eng.GET("/users/:id", usersHandlers.Show)
		req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID+1), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerCreateSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username"}
		defer func() {
			user.Read()
			user.Delete()
		}()
		h.eng.POST("/users", usersHandlers.Create)
		req := httptest.NewRequest("POST", "/users", createJsonParams(handlers.UserCraeteJSON{
			Name:     "username",
			Password: "password",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerCreateFailed(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username"}
		defer func() {
			user.Read()
			user.Delete()
		}()
		h.eng.POST("/users", usersHandlers.Create)
		req := httptest.NewRequest("POST", "/users", createJsonParams(handlers.UserCraeteJSON{
			Password: "password",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerUpdateSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username"}
		user.SetPassword("12345678")
		user.Create()
		defer user.Delete()
		h.eng.PATCH("/users/:id", usersHandlers.Update)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/users/%d", user.ID), createJsonParams(handlers.UserUpdateJSON{
			Name: "new_username",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
		user.Read()
		if user.Name != "new_username" {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerUpdateFailedNotFound(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.PATCH("/users/:id", usersHandlers.Update)
		req := httptest.NewRequest("PATCH", "/users/1", createJsonParams(handlers.UserUpdateJSON{
			Name: "new_username",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusNotFound != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerUpdateFailedInvalidParams(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.PATCH("/users/:id", usersHandlers.Update)
		req := httptest.NewRequest("PATCH", "/users/1", createJsonParams(handlers.UserUpdateJSON{}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusBadRequest != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerDeleteSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username", Password: "1234567890"}
		user.Create()
		defer user.Delete()
		h.eng.DELETE("/users/:id", usersHandlers.Delete)
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestUsersHandlerDeleteFailed(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "username", Password: "1234567890"}
		user.Create()
		defer user.Delete()
		h.eng.DELETE("/users/:id", usersHandlers.Delete)
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/users/%d", user.ID+1), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}
