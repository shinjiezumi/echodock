package board

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
)

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	flushMsg := util.GetFlushMsg(c)

	data := struct {
		Title    string
		Board    board.Board
		FlushMsg string
		Csrf     string
	}{
		Title:    util.GenerateTitle("掲示板一覧"),
		Board:    *b,
		FlushMsg: flushMsg,
		Csrf:     c.Get("csrf").(string),
	}

	return c.Render(http.StatusOK, "show", data)
}
