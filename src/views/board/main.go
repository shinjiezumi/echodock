package board

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func LoadTemplate(e *echo.Echo) {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/board/*.html")),
	}

	e.Renderer = t
}
