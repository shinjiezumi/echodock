package cookie

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func HandleCookie(c echo.Context) error {
	nameCookie, _ := readCookie(c)
	if nameCookie == nil {
		nameCookie = writeCookie(c)
	}
	return c.String(http.StatusOK, "name="+nameCookie.Name+", value="+nameCookie.Value)
}

func readCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie("name")
	if err != nil {
		return nil, err
	}

	log.Println(cookie.Name)
	log.Println(cookie.Value)
	return cookie, nil
}

func writeCookie(c echo.Context) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "name"
	cookie.Value = "john"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return cookie
}
