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

var postsHandlers = handlers.NewPostsHandler()

func TestPostsHandlerIndexSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		for i := 0; i < 10; i++ {
			post := models.Post{
				UserID:   user.ID,
				User:     user,
				Title:    fmt.Sprintf("title_%d", i),
				Problem:  "problem",
				Solution: "solution",
				Lesson:   "lesson",
			}
			post.Create()
			defer post.Delete()
		}

		login(h.eng, user.ID)

		h.eng.GET("/posts", postsHandlers.Index)
		req := httptest.NewRequest("GET", "/posts?page=2&limit=3&orderBy=title", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
		var resPosts models.Posts
		if err := json.Unmarshal(h.rec.Body.Bytes(), &resPosts); err != nil {
			t.Fatal("falied", h.rec)
		}
		if len(resPosts) != 3 {
			t.Fatal("falied", h.rec)
		}
		if resPosts[0].Title != "title_3" {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerIndex_NotLogin(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.GET("/posts", postsHandlers.Index)
		req := httptest.NewRequest("GET", "/posts?page=2&limit=3&orderBy=title", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}