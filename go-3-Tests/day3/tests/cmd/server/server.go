package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	// Aqui estamos descartando o erro
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	return r
}

func MakeRequest(
	method,
	url,
	body string,
) (
	*http.Request,
	*httptest.ResponseRecorder,
) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}
