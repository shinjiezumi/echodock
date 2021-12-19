package localstack

import (
	"github.com/labstack/echo/v4"

	"echodock/localstack/http/sqs"
)

func Setup(e *echo.Echo) {
	sqs.SetUpRoutes(e)
}
