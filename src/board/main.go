package board

import (
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {
	SetUpRoute(e)
}
