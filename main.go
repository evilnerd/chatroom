package main

import (
	"chatroom/data"
	"chatroom/handlers"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create main room
	room := data.NewRoom(data.MainRoom)
	data.AddRoom(room)
	room.Listen()

	// Set up the web server
	e := handlers.SetupMux()

	// Set up handling for ctrl-C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer close(sigs)

	go func() {
		<-sigs
		e.Logger.Printf("Exiting server.")
		e.Server.Close()
	}()

	// Run the server
	e.Logger.Fatal(e.Start(":8080"))
}
