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
				UserID: user.ID,
				Title:  fmt.Sprintf("title_%d", i),
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

func TestPostsHandlerShow_Success(t *testing.T) {
	readyServe(func(h *handlerTest) {
		post := models.Post{Title: "title"}
		post.Create()
		defer post.Delete()
		h.eng.GET("/post/:id", postsHandlers.Show)
		req := httptest.NewRequest("GET", fmt.Sprintf("/post/%d", post.ID), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerShow_NotFound(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.GET("/post/:id", postsHandlers.Show)
		req := httptest.NewRequest("GET", fmt.Sprintf("/post/%d", 1), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerCreate_Success(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		login(h.eng, user.ID)

		h.eng.POST("/posts", postsHandlers.Index)
		req := httptest.NewRequest("POST", "/posts", createJsonParams(handlers.PostCraeteJSON{
			Title: "title",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerCreate_NotLogin(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.POST("/posts", postsHandlers.Index)
		req := httptest.NewRequest("POST", "/posts", createJsonParams(handlers.PostCraeteJSON{
			Title: "title",
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerUpdate_Success(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		post := models.Post{
			Title:  "titile",
			UserID: user.ID,
		}
		post.Create()
		defer post.Delete()
		tag1 := models.Tag{Title: "a"}
		defer models.DB.Delete(&tag1)
		tag2 := models.Tag{Title: "b"}
		defer models.DB.Delete(&tag2)
		login(h.eng, user.ID)

		title := "new_title"
		tags := []string{"a", "b"}

		h.eng.PATCH("/posts/:id", postsHandlers.Update)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/posts/%d", post.ID), createJsonParams(handlers.PostUpdateJSON{
			Title: &title,
			Tags:  tags,
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
		post.Read()
		if post.Title != "new_title" {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerUpdate_NotFound(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		login(h.eng, user.ID)

		title := "new_title"

		h.eng.PATCH("/posts/:id", postsHandlers.Update)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/posts/%d", 1), createJsonParams(handlers.PostUpdateJSON{
			Title: &title,
		}))
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerDelete_Success(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		post := models.Post{
			Title:  "titile",
			UserID: user.ID,
		}
		post.Create()
		defer post.Delete()
		h.eng.DELETE("/posts/:id", postsHandlers.Delete)
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/posts/%d", post.ID), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestPostsHandlerDelete_Failed(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.DELETE("/post/:id", postsHandlers.Delete)
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/posts/%d", 1), nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK == h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}
