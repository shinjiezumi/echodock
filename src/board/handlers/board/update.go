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

// Update は掲示板を更新します
func Update(c echo.Context) error {
	req := new(StoreRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ederr.BadRequest
	}

	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return ederr.ResouceNotFound
	}

	tx := database.Conn.Begin()

	// 掲示板保存
	b.Name = req.Name
	b.Title = req.Title
	b.Body = req.Body
	board.SaveBoard(tx, b)

	// タグ削除
	board.DeleteTagRelation(tx, id)
	// タグ保存
	if len(req.Tags) > 0 {
		tr := make([]board.TagRelation, 0, len(req.Tags))
		for _, tag := range req.Tags {
			tr = append(tr, board.TagRelation{
				BoardID: b.ID,
				TagID:   tag,
			})
		}
		board.SaveTagRelation(tx, &tr)
	}

	if err = tx.Commit().Error; err != nil {
		panic(err)
	}

	util.SetFlushMsg(c, "更新しました")

	return c.Redirect(http.StatusFound, "/boards/"+strconv.Itoa(b.ID))
}
