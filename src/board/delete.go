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

// Delete は投稿を削除します
func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ederr.BadRequest
	}

	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return ederr.ResouceNotFound
	}

	tx := database.Conn.Begin()

	// 掲示板削除
	board.DeleteBoard(tx, id)

	// タグ保存
	if len(b.Tags) > 0 {
		board.DeleteTagRelation(tx, id)
	}

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}

	util.SetFlushMsg(c, "削除しました")

	_ = c.Redirect(http.StatusFound, "/boards")

	return nil
}
