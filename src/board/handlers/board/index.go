package board

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
)

type boardData struct {
	ID        int
	Title     string
	Body      string
	Name      string
	CreatedAt string
	UpdatedAt string
}

// Index は掲示板ページを表示します
func Index(c echo.Context) error {
	boardList := board.GetBoardList(database.Conn)

	boards := make([]boardData, len(boardList))
	for _, b := range boardList {
		boards = append(boards, boardData{
			ID:        b.ID,
			Title:     b.Title,
			Body:      b.Body,
			Name:      b.Name,
			CreatedAt: b.CreatedAt.Format(util.DateFormat),
			UpdatedAt: b.UpdatedAt.Format(util.DateFormat),
		})
	}

	flushMsg := util.GetFlushMsg(c)
	data := struct {
		Title    string
		Boards   []boardData
		FlushMsg string
		Csrf     string
	}{
		Title:    util.GeneratePageTitle("掲示板一覧"),
		Boards:   boards,
		FlushMsg: flushMsg,
		Csrf:     c.Get("csrf").(string),
	}

	return c.Render(http.StatusOK, "index", data)
}
