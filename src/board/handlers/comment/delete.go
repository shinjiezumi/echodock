package comment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"echodock/database"
	ederr "echodock/error"
	"echodock/models/board"
	"echodock/util"
)

// Delete はコメントを削除します
func Delete(c echo.Context) error {
	boardID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ederr.BadRequest
	}
	commentID, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return ederr.BadRequest
	}
	comment := board.GetComment(database.Conn, boardID, commentID)
	if comment == nil {
		return ederr.ResouceNotFound
	}

	// コメント削除
	board.DeleteComment(database.Conn, comment)

	util.SetFlushMsg(c, "コメントを削除しました")

	return c.Redirect(http.StatusFound, fmt.Sprintf("/boards/%d", boardID))
}
