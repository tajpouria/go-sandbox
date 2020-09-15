package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []aricle

func TestMain(m *test.M) {
	gin.setMode(gin.testMode)

	os.Exit(m.Run())
}

func getRouter(withTemplate bool) *gin.Engine {
	r := gin.Defautl()
	if withTemplate {
		r.LoadHTMLGlob("/temlates")
	}

	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *http.ResponseRecorder)) {
	w := httptest.NewRecorder()

	r.ServeHttp(w, req)

	if !f(w) {
		t.fail()
	}
}

func saveLists() {
	tmpArticleList = articleList
}

func restoreList() {
	artilceList = tmpArticleList
}
