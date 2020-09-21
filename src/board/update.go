package board

import (
	"github.com/labstack/echo/v4"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/models/board"
	"github.com/shinjiezumi/echodock/src/util"
	"net/http"
	"strconv"
)

func Update(c echo.Context) error {
	req := new(StoreRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	b := GetBoardByID(database.Conn, id)
	if b == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	tx := database.Conn.Begin()

	// 掲示板保存
	b.Name = req.Name
	b.Title = req.Title
	b.Body = req.Body
	SaveBoard(tx, b)

	// タグ削除
	DeleteTagRelation(tx, id)
	// タグ保存
	if len(req.Tags) > 0 {
		tr := make([]board.TagRelation, 0, len(req.Tags))
		for _, tag := range req.Tags {
			tr = append(tr, board.TagRelation{
				BoardID: b.ID,
				TagID:   tag,
			})
		}
		SaveTagRelation(tx, &tr)
	}

	tx.Commit()

	util.SetFlushMsg(c, "更新しました")

	c.Redirect(http.StatusFound, "/boards/"+strconv.Itoa(b.ID))

	return nil
}
