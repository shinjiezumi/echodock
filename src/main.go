package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echodock/board"
	"echodock/echobasic/cookie"
	"echodock/echobasic/request"
	"echodock/echobasic/response"
	"echodock/localstack/sqs"
	"echodock/views"
)

func main() {
	e := echo.New()
	views.LoadTemplate(e)

	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(".env"); err != nil {
			panic("load env file failed")
		}
	}

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

	// クエリパラメータの取得
	e.GET("/request/query", request.HandleQuery)

	// URLパラメータの取得
	e.GET("/request/get/:id", request.HandleGet)

	// POSTパラメータの取得
	e.POST("/request/post/:id", request.HandlePost)

	// 文字列レスポンス
	e.GET("/response/string", response.StringResponse)

	// HTMLレスポンス
	e.GET("/response/html", response.HtmlResponse)

	// HTML(テンプレートファイル)レスポンス
	e.GET("/response/htmltemplate", response.HtmlTemplateResponse)

	// Cookieの取得と設定
	e.GET("/cookie", cookie.HandleCookie)

	// 外部リソースの取得
	e.GET("/search/:query", func(c echo.Context) error {
		endpoint := "https://google.com/search?q=" + c.Param("query")
		req, _ := http.NewRequest("GET", endpoint, nil)

		var client = &http.Client{}
		resp, _ := client.Do(req)
		body, _ := ioutil.ReadAll(resp.Body)

		return c.HTML(http.StatusOK, string(body))
	})

	board.SetUpRoute(e)
	sqs.SetUp(e)

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
