package board

import (
	"github.com/labstack/echo/v4"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/models/board"
	"github.com/shinjiezumi/echodock/src/util"
	"net/http"
)

type StoreRequest struct {
	Name  string `form:"name" validate:"required,min=1,max=255"`
	Title string `form:"title" validate:"required,min=1,max=255"`
	Body  string `form:"body" validate:"required,min=1,max=1024"`
	Tags  []int  `form:"tags[]"`
}

func Store(c echo.Context) error {
	req := new(StoreRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	tx := database.Conn.Begin()

	// 掲示板保存
	b := board.Board{
		Name:  req.Name,
		Title: req.Title,
		Body:  req.Body,
	}
	board.SaveBoard(tx, &b)

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

	tx.Commit()

	util.SetFlushMsg(c, "作成しました")

	c.Redirect(http.StatusFound, "/boards")

	return nil
}
