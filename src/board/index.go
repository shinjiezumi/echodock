package board

import (
	"github.com/labstack/echo/v4"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/models/board"
	"github.com/shinjiezumi/echodock/src/util"
	"net/http"
)

type b struct {
	ID        int
	Title     string
	Body      string
	Name      string
	CreatedAt string
	UpdatedAt string
}

func Index(c echo.Context) error {
	boards := board.GetBoardList(database.Conn)

	var boardData []b
	for _, board := range boards {
		boardData = append(boardData, b{
			ID:        board.ID,
			Title:     board.Title,
			Body:      board.Body,
			Name:      board.Name,
			CreatedAt: board.CreatedAt.Format(util.DateFormat),
			UpdatedAt: board.UpdatedAt.Format(util.DateFormat),
		})
	}

	flushMsg := util.GetFlushMsg(c)
	data := struct {
		Title    string
		Boards   []b
		FlushMsg string
		Csrf     string
	}{
		Title:    util.GenerateTitle("掲示板一覧"),
		Boards:   boardData,
		FlushMsg: flushMsg,
		Csrf:     c.Get("csrf").(string),
	}

	return c.Render(http.StatusOK, "index", data)
}
