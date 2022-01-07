package board

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echodock/database"
	ederr "echodock/error"
	"echodock/models/board"
	"echodock/util"
)

// Edit は掲示板編集ページを表示します
func Edit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ederr.BadRequest
	}
	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return ederr.ResouceNotFound
	}

	form := make(map[string]interface{}, 0)
	form["action"] = "/boards/" + strconv.Itoa(b.ID)
	form["method"] = http.MethodPost
	form["name"] = b.Name
	form["title"] = b.Title
	form["body"] = b.Body

	tags := board.GetTags(database.Conn)
	tagData := make([]tag, 0, len(tags))
	for _, t := range tags {
		isChecked := false
		for _, boardTag := range b.Tags {
			if t.ID == boardTag.ID {
				isChecked = true
			}
		}
		tagData = append(tagData, tag{
			ID:        t.ID,
			Name:      t.Name,
			IsChecked: isChecked,
		})
	}
	form["tags"] = tagData
	form["csrf"] = c.Get("csrf").(string)

	data := struct {
		Title string
		Form  map[string]interface{}
	}{
		Title: util.GeneratePageTitle("掲示板更新"),
		Form:  form,
	}

	return c.Render(http.StatusOK, "edit", data)
}
