package service

import (
	"chatroom/data"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func HasUser(c echo.Context) bool {
	return GetUser(c).Name != ""
}

func GetUser(c echo.Context) data.User {
	sess, err := session.Get(SessionName, c)
	if err != nil {
		c.Logger().Errorf("Checking for logged in user, but couldn't get to the whole session: %v\n", err)
		return data.User{}
	}

	user := sess.Values[UserKey]
	if user == nil {
		return data.User{}
	}
	return user.(data.User)
}
