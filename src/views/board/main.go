package board

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

const basePath = "views/board/"

var pages = []string{
	"index",
	"create",
	"show",
	"edit",
}

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates[name].ExecuteTemplate(w, "layout.html", data)
}

func LoadTemplate(e *echo.Echo) {
	commonTemplates := []string{
		basePath + "layout.html",
		basePath + "header.html",
		basePath + "form.html",
	}

	templates := make(map[string]*template.Template)
	for _, v := range pages {
		templates[v] = template.Must(template.ParseFiles(append(commonTemplates, basePath+v+".html")...))
	}

	t := &Template{
		templates: templates,
	}

	e.Renderer = t
}
