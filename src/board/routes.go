package board

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/views/board"
)

func SetUpRoute(e *echo.Echo) {
	database.Initialize()

	board.LoadTemplate(e)
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.Static("/", "assets/board")

	e.GET("/boards", Index)
	e.GET("/boards/create", Create)
	e.POST("/boards/create", Store)
}
