package main

import (
	"fmt"

	"golang.org/x/net/websocket"
)

// Global variable to keep track of WebSocket connections
var connections []*websocket.Conn

// WebSocket handler to add clients and remove them when disconnected
func websocketHandler(ws *websocket.Conn) {
	defer ws.Close()

	// Add the new connection to the connections slice
	connections = append(connections, ws)
	fmt.Println("New WebSocket client connected")

	// Remove the connection from the slice when disconnected
	defer func() {
		for i, conn := range connections {
			if conn == ws {
				// Remove the connection at index i from the slice
				connections = append(connections[:i], connections[i+1:]...)
				fmt.Println("WebSocket client disconnected")
				break
			}
		}
	}()

	// Keep the WebSocket connection open and listen for messages
	for {
		var msg string

		// Wait for the message from the client
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			fmt.Println("WebSocket connection closed", err)
			break
		}
	}
}

// Broadcast Twitch messages to all WebSocket clients
func broadcastToWebsocketClients(message string) {
	for _, conn := range connections {
		if err := websocket.Message.Send(conn, message); err != nil {
			fmt.Println("Error sending WebSocket message:", err)
		}
	}
}
