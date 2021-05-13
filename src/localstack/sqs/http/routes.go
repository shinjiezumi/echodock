package http

import "github.com/labstack/echo/v4"

const prefix = "/localstack/sqs"

func SetUpRoutes(e *echo.Echo) {
	r := e.Group(prefix)

	r.GET("/queues", GetQueues)
}
