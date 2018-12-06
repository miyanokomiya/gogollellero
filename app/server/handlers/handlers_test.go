package handlers_test

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/miyanokomiya/gogollellero/app/server"
	"github.com/miyanokomiya/gogollellero/app/server/constants"
	"github.com/miyanokomiya/gogollellero/app/server/handlers"
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

func mockPost(uri string, params interface{}) *httptest.ResponseRecorder {
	var reader *strings.Reader
	if params != nil {
		json, _ := json.Marshal(params)
		reader = strings.NewReader(string(json))
	}
	req := httptest.NewRequest("POST", uri, reader)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router := server.Create()
	router.ServeHTTP(rec, req)
	return rec
}

func mockGet(uri string, params url.Values) *httptest.ResponseRecorder {
	return mockQuery("GET", uri, params)
}

func mockDelete(uri string, params url.Values) *httptest.ResponseRecorder {
	return mockQuery("DELETE", uri, params)
}

func mockQuery(method string, uri string, params url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, uri, nil)
	if params != nil {
		req.URL.RawQuery = params.Encode()
	}
	rec := httptest.NewRecorder()
	router := server.Create()
	router.ServeHTTP(rec, req)
	return rec
}

func authWrapper(fn func(user *models.User)) {
	models.GormOpen()
	defer models.GormClose()
	user := models.User{Name: "username"}
	user.SetPassword("1234567890")
	user.Create()
	defer user.Delete()

	json := handlers.LoginJson{
		UserName: user.Name,
		Password: "1234567890",
	}
	mockPost("/api/v1/login", json)

	fn(&user)
}
