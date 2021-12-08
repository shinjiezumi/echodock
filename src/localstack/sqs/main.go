package sqs

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echodock/database"
	"echodock/localstack/sqs/http"
)

func SetUp(e *echo.Echo) {
	database.Initialize()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	http.SetUpRoutes(e)
}
