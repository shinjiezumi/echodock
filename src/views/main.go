package views

import (
	"html/template"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"

	_ "echodock/views/statik"
)

const (
	basePath    = "/"
	boardPath   = "/board/"
	commentPath = "/board/comment/"
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
		t := template.New("t").Funcs(funcs)
		filenames := append(commonTemplates, basePath+v+".html")
		templates[v] = parseFiles(t, filenames)
		templates[v].Funcs(funcs)
	}
	for _, v := range boardPages {
		t := template.New("t").Funcs(funcs)
		filenames := append(commonTemplates, boardPath+v+".html")
		templates[v] = parseFiles(t, filenames)
		templates[v].Funcs(funcs)
	}

	t := &Template{
		templates: templates,
	}

	e.Renderer = t
}

func parseFiles(t *template.Template, filenames []string) *template.Template {
	statikFS, err := fs.NewWithNamespace("views")
	if err != nil {
		panic(err)
	}

	// template.ParseFiles流用
	for _, filename := range filenames {
		f, err := statikFS.Open(filename)
		if err != nil {
			panic(err)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		s := string(b)
		name := filepath.Base(filename)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			panic(err)
		}
	}
	return t

}
