package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gempir/go-twitch-irc/v3"
	"golang.org/x/net/websocket"
)

// go run *.go
func main() {
	// Initialize the Twitch client
	client := twitch.NewAnonymousClient()

	// Set up the Twitch handlers
	twitchHandler(client)

	// WebSocket server setup
	http.Handle("/ws", websocket.Handler(websocketHandler))
	fmt.Println("WebSocket server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
