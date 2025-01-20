package main

import (
	"github.com/eric-ming2/hanabi/hanabi-frontend/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
)

func listen(conn *websocket.Conn, wsResChan chan WorkerResponse) {
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
			var response generated.Response
			err := proto.Unmarshal(message, &response)
			if err != nil {
				log.Printf("Unable to unmarshal proto: %v", err)
				return
			}
			switch response.ResponseType {
			case generated.ResponseType_UPDATE_GAME:
				log.Printf("Received UPDATE_GAME message: %v", response.GetUpdateGame())
				wsResChan <- WorkerResponse{
					Type:    UpdateGameState,
					Payload: response.GetUpdateGame(),
				}
			}
		default:
			log.Printf("Received unknown message type (%d): %x", messageType, message)
		}
	}
}
