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

type StoreCommentRequest struct {
	Name    string `form:"name" validate:"required,min=1,max=255"`
	Comment string `form:"comment" validate:"required,min=1,max=1024"`
}

func Store(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := board.GetBoardByID(database.Conn, id)
	if b == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	req := new(StoreCommentRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	tx := database.Conn.Begin()

	// コメント保存
	comment := board.Comment{
		BoardID: id,
		Name:    req.Name,
		Comment: req.Comment,
	}
	board.SaveComment(tx, &comment)

	tx.Commit()

	util.SetFlushMsg(c, "コメントしました")

	c.Redirect(http.StatusFound, fmt.Sprintf("/boards/%d", id))

	return nil
}
