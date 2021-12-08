package comment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
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
