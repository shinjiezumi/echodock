package sqs

import "github.com/labstack/echo/v4"

const prefix = "/localstack/sqs"

// SetUpRoutes はsqs関連のルーティングをセットアップします
func SetUpRoutes(e *echo.Echo) {
	r := e.Group(prefix)

	r.GET("/queues", GetQueues)
}
