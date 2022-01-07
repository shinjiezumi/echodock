package board

import (
	"net/http"
	"strconv"

	ederr "echodock/error"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
)

// Show は掲示板詳細ページを表示します
func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ederr.BadRequest
	}
	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return ederr.ResouceNotFound
	}

	flushMsg := util.GetFlushMsg(c)

	data := struct {
		Title    string
		Board    board.Board
		FlushMsg string
		Csrf     string
	}{
		Title:    util.GeneratePageTitle("掲示板一覧"),
		Board:    *b,
		FlushMsg: flushMsg,
		Csrf:     c.Get("csrf").(string),
	}

	return c.Render(http.StatusOK, "show", data)
}
