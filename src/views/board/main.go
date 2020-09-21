package board

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"strings"
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

	funcs := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}

	templates := make(map[string]*template.Template)
	for _, v := range pages {
		templates[v] = template.Must(template.New("t").Funcs(funcs).ParseFiles(append(commonTemplates, basePath+v+".html")...))
		templates[v].Funcs(funcs)
	}

	t := &Template{
		templates: templates,
	}

	e.Renderer = t
}
