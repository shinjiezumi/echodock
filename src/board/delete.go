package board

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
)

func Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	tx := database.Conn.Begin()

	// 掲示板削除
	board.DeleteBoard(tx, id)

	// タグ保存
	if len(b.Tags) > 0 {
		board.DeleteTagRelation(tx, id)
	}

	tx.Commit()

	util.SetFlushMsg(c, "削除しました")

	c.Redirect(http.StatusFound, "/boards")

	return nil
}
