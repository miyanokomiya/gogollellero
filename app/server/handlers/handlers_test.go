package handlers_test

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/miyanokomiya/gogollellero/app/server"
)

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
