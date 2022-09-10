package service

import (
	"chatroom/data"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	sockets []*websocket.Conn
)

func PublishMessagesSocket(upgrader websocket.Upgrader) func(e echo.Context) error {
	return func(c echo.Context) error {
		// Get the user
		user := GetUser(c)

		// Get the room
		roomName := c.Param("room")
		if roomName == "" {
			roomName = data.MainRoom
		}
		room := data.GetRoomByName(roomName)
		if room == nil {
			return fmt.Errorf("could not find room %s", roomName)
		}

		// Set up the websocket
		var close chan bool
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		ws.SetCloseHandler(func(code int, text string) error {
			// unregister
			room.UnregisterSocket(ws)
			close <- true
			return nil
		})

		// register
		room.RegisterSocket(ws)

		// this goroutine listens to my user sending a chat message.
		go func(shouldClose chan bool) {
			for {
				select {
				case c := <-shouldClose:
					if c {
						return
					}

				default:
					msg := data.SendMessage{}
					err = ws.ReadJSON(&msg)
					if err != nil {
						if websocket.IsCloseError(err) {
							c.Logger().Printf("Closing connection for user '%s'\n", user.Name)
						} else {
							c.Logger().Errorf("Error parsing incoming message for user '%s': %v\n", user.Name, err)
						}
					} else {
						room.Channel <- data.NewReceivedMessage(user, msg.Body)
					}
				}
			}
		}(close)

		return err
	}
}
