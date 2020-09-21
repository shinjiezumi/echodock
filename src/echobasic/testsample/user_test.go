package testsample

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mockDB = map[string]*User{
		"taro@gmail.com": {"Taro", "taro@gmail.com"},
	}
	userJSON = `{"name":"Taro","email":"taro@gmail.com"}`
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, strings.Replace(rec.Body.String(), "\n", "", -1))
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("taro@gmail.com")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, strings.Replace(rec.Body.String(), "\n", "", -1))
	}
}
