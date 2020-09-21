package board

import (
	"github.com/labstack/echo/v4"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/util"
	"net/http"
	"strconv"
)

func Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := GetBoardByID(database.Conn, id)
	if b == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	tx := database.Conn.Begin()

	// 掲示板削除
	DeleteBoard(tx, id)

	// タグ保存
	if len(b.Tags) > 0 {
		DeleteTagRelation(tx, id)
	}

	tx.Commit()

	util.SetFlushMsg(c, "削除しました")

	c.Redirect(http.StatusFound, "/boards")

	return nil
}
