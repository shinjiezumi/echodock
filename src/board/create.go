package board

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
)

type tag struct {
	ID        int
	Name      string
	IsChecked bool
}

func Create(c echo.Context) error {
	form := make(map[string]interface{}, 0)
	form["action"] = "/boards"
	form["method"] = "POST"
	form["name"] = ""
	form["title"] = ""
	form["body"] = ""

	tagData := make([]tag, 0)
	tags := board.GetTags(database.Conn)
	for _, t := range tags {
		tagData = append(tagData, tag{
			ID:   t.ID,
			Name: t.Name,
		})
	}
	form["tags"] = tagData
	form["csrf"] = c.Get("csrf").(string)

	data := struct {
		Title string
		Form  map[string]interface{}
	}{
		Title: util.GenerateTitle("掲示板作成"),
		Form:  form,
	}

	return c.Render(http.StatusOK, "create", data)
}
