package board

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"

	_ "echodock/assets/statik"
	"echodock/board/handlers/board"
	"echodock/board/handlers/comment"
)

// SetUpRoute はルーティングを設定します
func SetUpRoute(e *echo.Echo) {
	e.GET("/boards", board.Index)
	e.GET("/boards/create", board.Create)
	e.POST("/boards", board.Store)
	e.GET("/boards/:id", board.Show)
	e.GET("/boards/:id/edit", board.Edit)
	e.PUT("/boards/:id", board.Update)
	e.DELETE("/boards/:id", board.Delete)

	e.POST("/boards/:id/comments", comment.Store)
	e.DELETE("/boards/:id/comments/:comment_id", comment.Delete)

	statikFS, err := fs.NewWithNamespace("assets")
	if err != nil {
		panic(err)
	}
	assetHandler := http.FileServer(statikFS)
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets", assetHandler)))
}
