package request

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type LoginParam struct {
	Email    string `json:"email" format:"email"`
	Password string `json:"password" format:"password"`
}

func HandleQuery(c echo.Context) error {
	queryParams := c.QueryParams()
	var response string
	for k, v := range queryParams {
		response += k + ":" + v[0]
	}
	paramC := c.QueryParam("c")
	fmt.Println(paramC)
	return c.String(http.StatusOK, response)
}

func HandleGet(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id="+id)
}

func HandlePost(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var loginParam LoginParam
	if err := c.Bind(&loginParam); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(loginParam)
	return c.String(http.StatusOK, "email="+email+", password="+password)
}
