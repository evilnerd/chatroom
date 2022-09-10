package handlers

import (
	"chatroom/data"
	"chatroom/service"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func LogoutHandler(c echo.Context) error {
	sess, err := session.Get(service.SessionName, c)
	if err != nil {
		c.Logger().Errorf("Going to log out but can't get the whole session: %v\n", err)
		return err
	}
	sess.Values[service.UserKey] = nil
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		c.Logger().Errorf("Could not save the session after resetting the user: %v\n", err)
		return err
	}
	c.Redirect(http.StatusFound, "/")
	return nil

}

func LoginHandler(c echo.Context) error {
	sess, err := session.Get(service.SessionName, c)
	if err != nil {
		return fmt.Errorf("trying to register user but couldn't get to the whole session: %v", err)
	}

	name := c.Request().FormValue("userName")
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("empty name")
	}
	user := data.NewUser(name)
	sess.Values[service.UserKey] = user
	data.GetRoomByName(data.MainRoom).AnnounceNewUser(user)
	sess.Options = &sessions.Options{
		Path: "/",
	}
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		c.Logger().Errorf("Error saving the session: %v", err)
	}
	return c.Redirect(http.StatusFound, "/")
}
