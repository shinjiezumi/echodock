package echobasic

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	"echodock/echobasic/cookie"
	"echodock/echobasic/request"
	"echodock/echobasic/response"
)

func SetUpRoutes(e *echo.Echo) {
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
}
