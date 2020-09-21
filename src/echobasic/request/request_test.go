package request

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleQuery_WithNoQuery(t *testing.T) {
	e := echo.New()
	e.GET("/request/query", HandleQuery)

	req := httptest.NewRequest(http.MethodGet, "/request/query", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}

func TestHandleQuery_WithQuery(t *testing.T) {
	e := echo.New()
	e.GET("/request/query", HandleQuery)

	req := httptest.NewRequest(http.MethodGet, "/request/query?a=1&b=1", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	response := rec.Body.String()
	if response != "a:1b:1" && response != "b:1a:1" {
		t.Error("Invalid response")
	}
}
