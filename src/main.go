package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shinjiezumi/echodock/src/board"
	"github.com/shinjiezumi/echodock/src/echobasic/cookie"
	"github.com/shinjiezumi/echodock/src/echobasic/request"
	"github.com/shinjiezumi/echodock/src/echobasic/response"
	"github.com/shinjiezumi/echodock/src/views"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()
	views.LoadTemplate(e)

	// アクセスログの設定
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	// ROOT
	e.GET("/", func(c echo.Context) error {
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

	e.Logger.Fatal(e.Start(":8080"))
}
