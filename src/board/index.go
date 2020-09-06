package board

import (
	"github.com/labstack/echo"
	"net/http"
)

func Index(c echo.Context) error {
	data := struct {
		Title string
	}{
		Title: "テンプレート",
	}

	return c.Render(http.StatusOK, "index", data)
}
