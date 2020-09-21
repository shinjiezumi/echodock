package board

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinjiezumi/echodock/src/database"
	"github.com/shinjiezumi/echodock/src/views/board"
)

func SetUpRoute(e *echo.Echo) {
	database.Initialize()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	board.LoadTemplate(e)
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.Static("/", "assets/board")

	e.GET("/boards", Index)
	e.GET("/boards/create", Create)
	e.POST("/boards/create", Store)
}
