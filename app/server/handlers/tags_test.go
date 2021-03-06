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

var tagsHandlers = handlers.NewTagsHandler()

func TestTagsHandlerIndexSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		for i := 0; i < 20; i++ {
			tag := models.Tag{
				Title: fmt.Sprintf("title_%d", i),
			}
			models.DB.Create(&tag)
			defer models.DB.Delete(&tag)
		}

		h.eng.GET("/tags", tagsHandlers.Index)
		req := httptest.NewRequest("GET", "/tags?page=2&limit=3&orderBy=title&keyword=1", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
		var resTags models.Tags
		if err := json.Unmarshal(h.rec.Body.Bytes(), &resTags); err != nil {
			t.Fatal("falied", h.rec)
		}
		if len(resTags) != 3 {
			t.Fatal("falied", h.rec)
		}
		if resTags[0].Title != "title_12" {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestTagsHandlerIndexOfMineSuccess(t *testing.T) {
	readyServe(func(h *handlerTest) {
		user := models.User{Name: "user", Password: "password"}
		user.Create()
		defer user.Delete()
		var tags models.Tags
		for i := 0; i < 20; i++ {
			tag := models.Tag{
				Title: fmt.Sprintf("title_%d", i),
			}
			models.DB.Create(&tag)
			defer models.DB.Delete(&tag)
			if i < 3 {
				tags = append(tags, tag)
			}
		}
		post := models.Post{Title: "title"}
		post.UserID = user.ID
		post.Tags = tags
		post.Create()
		defer post.Delete()

		login(h.eng, user.ID)

		h.eng.GET("/tags", tagsHandlers.IndexOfMine)
		req := httptest.NewRequest("GET", "/tags", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusOK != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
		var resTags models.Tags
		if err := json.Unmarshal(h.rec.Body.Bytes(), &resTags); err != nil {
			t.Fatal("falied", h.rec)
		}
		if len(resTags) != 3 {
			t.Fatal("falied", h.rec)
		}
	})
}

func TestTagsHandlerIndexOfMine_NotLogin(t *testing.T) {
	readyServe(func(h *handlerTest) {
		h.eng.GET("/tags", tagsHandlers.IndexOfMine)
		req := httptest.NewRequest("GET", "/tags", nil)
		h.eng.ServeHTTP(h.rec, req)

		if http.StatusUnauthorized != h.rec.Code {
			t.Fatal("falied", h.rec)
		}
	})
}
