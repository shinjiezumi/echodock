package response

import (
	"github.com/labstack/echo"
	"net/http"
)

func StringResponse(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func HtmlResponse(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>HelloWorld</h1>")
}

func HtmlTemplateResponse(c echo.Context) error {
	data := struct {
		Title string
	}{
		Title: "テンプレート",
	}

	return c.Render(http.StatusOK, "hello", data)
}
