package views

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"strings"
)

const (
	basePath    = "views/"
	boardPath   = "views/board/"
	commentPath = "views/board/comment/"
)

var pages = []string{
	"top",
	"hello",
}

var boardPages = []string{
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
		boardPath + "form.html",
		commentPath + "comment.html",
		commentPath + "form.html",
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
	for _, v := range boardPages {
		templates[v] = template.Must(template.New("t").Funcs(funcs).ParseFiles(append(commonTemplates, boardPath+v+".html")...))
		templates[v].Funcs(funcs)
	}

	t := &Template{
		templates: templates,
	}

	e.Renderer = t
}
