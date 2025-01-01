package main

import (
	"fmt"
	"log"

	"github.com/gempir/go-twitch-irc/v3"
)

func twitchHandler(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		formattedMessage := fmt.Sprintf("[%s] %s", message.User.DisplayName, message.Message)

		// Log the message to console
		fmt.Println(formattedMessage)

		// Broadcast the message to all WebSocket clients
		broadcastToWebsocketClients(formattedMessage)
	})

	client.OnConnect(func() {
		fmt.Println("Connected to Twitch chat!")

		// Channel to join aka twitch streamer name
		client.Join("")
	})

	go func() {
		// Start the Twitch client
		if err := client.Connect(); err != nil {
			log.Fatal("Error connecting to Twitch:", err)
		}
	}()
}
