package board

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shinjiezumi/echodock/src/board/comment"
	"github.com/shinjiezumi/echodock/src/database"
)

func SetUpRoute(e *echo.Echo) {
	database.Initialize()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))
	e.Static("/", "assets/board")

	e.GET("/boards", Index)
	e.GET("/boards/create", Create)
	e.POST("/boards", Store)
	e.GET("/boards/:id", Show)
	e.GET("/boards/:id/edit", Edit)
	e.PUT("/boards/:id", Update)
	e.DELETE("/boards/:id", Delete)

	e.POST("/boards/:id/comments", comment.Store)
	e.DELETE("/boards/:id/comments/:comment_id", comment.Delete)
}
