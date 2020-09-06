package board

import (
	"github.com/labstack/echo"
	"github.com/shinjiezumi/echodock/src/views/board"
)

func SetUpRoute(e *echo.Echo) {
	board.LoadTemplate(e)

	e.Static("/", "assets/board")

	e.GET("/boards", Index)
}
