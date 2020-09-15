package board

import (
	"github.com/labstack/echo"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/views/board"
)

func SetUpRoute(e *echo.Echo) {
	database.Initialize()

	board.LoadTemplate(e)

	e.Static("/", "assets/board")

	e.GET("/boards", Index)
}
