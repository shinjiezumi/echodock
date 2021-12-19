package comment

import (
	"fmt"
	"net/http"
	"strconv"

	ederr "echodock/error"

	"github.com/labstack/echo/v4"

	"echodock/database"
	"echodock/models/board"
	"echodock/util"
)

// StoreCommentRequest コメント保存リクエスト
type StoreCommentRequest struct {
	Name    string `form:"name" validate:"required,min=1,max=255"`     // コメント投稿者
	Comment string `form:"comment" validate:"required,min=1,max=1024"` // コメント
}

// Store はコメントを保存します
func Store(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ederr.BadRequest
	}

	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return ederr.ResouceNotFound
	}

	req := new(StoreCommentRequest)
	if err = c.Bind(req); err != nil {
		panic(err)
	}

	// コメント保存
	tx := database.Conn.Begin()

	comment := board.Comment{
		BoardID: id,
		Name:    req.Name,
		Comment: req.Comment,
	}
	board.SaveComment(tx, &comment)

	if err = tx.Commit().Error; err != nil {
		panic(err)
	}

	util.SetFlushMsg(c, "コメントしました")

	return c.Redirect(http.StatusFound, fmt.Sprintf("/boards/%d", id))
}
