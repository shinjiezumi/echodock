package util

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func SetFlushMsg(c echo.Context, msg string) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["flushMsg"] = msg
	sess.Save(c.Request(), c.Response())
}

func GetFlushMsg(c echo.Context) string {
	sess, err := session.Get("session", c)
	if err != nil {
		return ""
	}

	if msg, ok := sess.Values["flushMsg"]; ok {
		delete(sess.Values, "flushMsg")
		sess.Save(c.Request(), c.Response())
		return msg.(string)
	}

	return ""
}
