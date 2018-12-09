package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server/constants"
	"github.com/miyanokomiya/gogollellero/app/server/models"
)

type handlerTest struct {
	rec *httptest.ResponseRecorder
	eng *gin.Engine
	ctx *gin.Context
}

func readyServe(fn func(h *handlerTest)) {
	models.GormOpen()
	defer models.GormClose()

	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	ctx, eng := gin.CreateTestContext(rec)
	store := cookie.NewStore([]byte("secret"))
	eng.Use(sessions.Sessions(constants.CookieName, store))
	fn(&handlerTest{rec: rec, eng: eng, ctx: ctx})
}

func login(eng *gin.Engine, userID int) {
	eng.Use(func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Set(constants.SessionUser, userID)
		if err := session.Save(); err != nil {
			panic(err)
		}
		ctx.Next()
	})
}

func createJsonParams(params interface{}) *strings.Reader {
	json, _ := json.Marshal(params)
	return strings.NewReader(string(json))
}

func userListWrapper(count int, fn func(models.Users)) {
	var users models.Users
	for i := 0; i < count; i++ {
		user := models.User{Name: fmt.Sprintf("user_%d", i), Password: "abcdabcd"}
		user.Create()
		defer user.Delete()
		users = append(users, user)
	}
	fn(users)
}

func postListWrapper(count int, fn func(models.Posts)) {
	var posts models.Posts
	for i := 0; i < count; i++ {
		user := models.User{Name: fmt.Sprintf("user_%d", i), Password: "abcdabcd"}
		user.Create()
		defer user.Delete()
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
		posts = append(posts, post)
	}
	fn(posts)
}
