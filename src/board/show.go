package board

import (
	"github.com/labstack/echo/v4"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/models/board"
	"github.com/shinjiezumi/echodock/src/util"
	"net/http"
	"strconv"
)

func Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := GetBoardByID(database.Conn, id)
	if b == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	flushMsg := util.GetFlushMsg(c)

	data := struct {
		Title    string
		Board    board.Board
		FlushMsg string
	}{
		Title:    util.GenerateTitle("掲示板一覧"),
		Board:    *b,
		FlushMsg: flushMsg,
	}

	return c.Render(http.StatusOK, "show", data)
}
