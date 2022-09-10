package handlers

import (
	"chatroom/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	if service.HasUser(c) {
		return c.Render(http.StatusOK, "index.html", nil)
	} else {
		return c.Redirect(http.StatusFound, "/login")
	}
}
