package error

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var BadRequest = echo.NewHTTPError(http.StatusBadRequest, MessageBadRequest)
var ResouceNotFound = echo.NewHTTPError(http.StatusNotFound, MessageResourceNotFound)
