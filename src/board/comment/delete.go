package comment

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/models/board"
	"github.com/shinjiezumi/echodock/src/util"
	"net/http"
	"strconv"
)

func Delete(c echo.Context) error {
	boardID, _ := strconv.Atoi(c.Param("id"))
	commentID, _ := strconv.Atoi(c.Param("comment_id"))
	comment := board.GetComment(database.Conn, boardID, commentID)
	if comment == nil {
		return echo.NewHTTPError(http.StatusNotFound, "comment not found")
	}

	// コメント削除
	board.DeleteComment(database.Conn, boardID, commentID)

	util.SetFlushMsg(c, "コメントを削除しました")

	c.Redirect(http.StatusFound, fmt.Sprintf("/boards/%d", boardID))

	return nil
}
