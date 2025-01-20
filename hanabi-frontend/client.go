package main

import (
	"github.com/eric-ming2/hanabi/hanabi-frontend/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/url"
)

func connect() {
	serverURL := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:8080",
		Path:   "/",
	}

	log.Printf("Connecting to WebSocket server at %s", serverURL.String())

	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()
	log.Println("Connected to WebSocket server")

	binaryData, err := proto.Marshal(createStartGameRequest())
	if err != nil {
		log.Fatalf("Failed to marshal proto", err)
	}

	err = conn.WriteMessage(websocket.BinaryMessage, binaryData)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// Handle different message types
		switch messageType {
		case websocket.TextMessage:
			log.Printf("Received text message: %s", message)
		case websocket.BinaryMessage:
			// TODO: Your naming sucks, this should be Response, not Request
			var request generated.Request
			err := proto.Unmarshal(message, &request)
			if err != nil {
				log.Printf("Unable to unmarshal proto: %v", err)
				return
			}
			switch request.RequestType {
			case generated.RequestType_UPDATE_GAME:
				log.Printf("Received UPDATE_GAME message: %v", request.GetUpdateGame())
			}
		default:
			log.Printf("Received unknown message type (%d): %x", messageType, message)
		}
	}
}
func createStartGameRequest() *generated.Request {
	startGameRequest := &generated.StartGameRequest{}
	return &generated.Request{
		RequestType: generated.RequestType_START_GAME,
		Request: &generated.Request_StartGame{
			StartGame: startGameRequest,
		},
	}
}
