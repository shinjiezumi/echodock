package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echodock/board"
	"echodock/database"
	"echodock/echobasic"
	"echodock/localstack"
	"echodock/views"
)

func main() {
	e := echo.New()
	views.LoadTemplate(e)

	if os.Getenv("APP_ENV") == "local" {
		// 環境変数読み込み
		if err := godotenv.Load(".env"); err != nil {
			panic("load env file failed")
		}
	}

	// DB初期化
	if err := database.Initialize(); err != nil {
		panic(fmt.Sprintf("db connection can't not initialized! err: %s", err))
	}

	// Session
	if secret := os.Getenv("SESSION_KEY"); secret == "" {
		panic("SESSION_KEY is empty")
	} else {
		e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	}
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	// CSRF
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	// アクセスログの設定
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.Recover())

	// ROOT
	e.GET("/", top)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	// echo基本機能
	echobasic.SetUpRoutes(e)

	// 掲示板
	board.Setup(e)

	// localstack系
	localstack.Setup(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func top(c echo.Context) error {
	data := struct {
		Title string
	}{
		Title: "echodock",
	}

	return c.Render(http.StatusOK, "top", data)
}
