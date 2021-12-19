package board

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"

	_ "echodock/assets/statik"
	"echodock/board/comment"
)

// SetUpRoute はルーティングを設定します
func SetUpRoute(e *echo.Echo) {
	e.GET("/boards", Index)
	e.GET("/boards/create", Create)
	e.POST("/boards", Store)
	e.GET("/boards/:id", Show)
	e.GET("/boards/:id/edit", Edit)
	e.PUT("/boards/:id", Update)
	e.DELETE("/boards/:id", Delete)

	e.POST("/boards/:id/comments", comment.Store)
	e.DELETE("/boards/:id/comments/:comment_id", comment.Delete)

	statikFS, err := fs.NewWithNamespace("assets")
	if err != nil {
		panic(err)
	}
	assetHandler := http.FileServer(statikFS)
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets", assetHandler)))
}
