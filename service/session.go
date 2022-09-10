package service

import (
	"chatroom/data"
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SessionSecret = "blablahsecret"
	SessionName   = "chatroom-session"
	UserKey       = "userName"
)

func SetupSession(e *echo.Echo) {
	e.Logger.Print("Setting up sessions\n")
	gob.Register(data.User{})
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SessionSecret))))
}
