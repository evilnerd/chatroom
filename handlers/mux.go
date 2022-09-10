package handlers

import (
	"chatroom/service"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func SetupMux() *echo.Echo {
	e := echo.New()
	setupMiddlewares(e)
	setupRoutes(e)
	setupTemplates(e)
	setupCors(e)
	setupMetrics(e)
	return e
}

func setupCors(e *echo.Echo) {
	e.Use(middleware.CORS())
}

func setupMetrics(e *echo.Echo) {
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
}

func setupTemplates(e *echo.Echo) {
	e.Renderer = service.NewTemplateRenderer("./templates/*.html")
}

func setupMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	service.SetupSession(e)
}

func setupRoutes(e *echo.Echo) {
	// Static
	e.Static("/public/", "public")

	// Main
	e.GET("/", Home)
	e.GET("/logout", LogoutHandler)

	// Login
	loginGroup := e.Group("/login")
	loginGroup.GET("", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	})
	loginGroup.POST("", LoginHandler)

	// Messages

	e.GET("/ws", service.PublishMessagesSocket(websocket.Upgrader{}))
	e.GET("/ws/:room", service.PublishMessagesSocket(websocket.Upgrader{}))
}
