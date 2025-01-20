package main

import (
	"github.com/eric-ming2/hanabi/hanabi-frontend/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/url"
)

func clientWorker(workerReqChan chan WorkerRequest, workerResChan chan WorkerResponse) {
	for req := range workerReqChan {
		switch req.Type {
		case ConnectRequest:
			payload, ok := req.Payload.(ConnectRequestPayload)
			if !ok {
				log.Fatalf("Failed to cast payload to ConnectRequestPayload")
			}
			connect(workerResChan, payload)
		}
	}
}

func connect(workerResChan chan WorkerResponse, payload ConnectRequestPayload) {
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

	initConnectionReq, err := proto.Marshal(createInitConnectionRequest(payload))
	if err != nil {
		log.Fatalf("Failed to marshal proto", err)
	}
	err = conn.WriteMessage(websocket.BinaryMessage, initConnectionReq)
	if err != nil {
		log.Fatalf("Failed to write init connection message", err)
	}

	startGameReq, err := proto.Marshal(createStartGameRequest())
	if err != nil {
		log.Fatalf("Failed to marshal proto", err)
	}

	// TODO: Move into sender fn, send on button click
	err = conn.WriteMessage(websocket.BinaryMessage, startGameReq)
	if err != nil {
		log.Fatalf("Failed to write start game message", err)
	}

	go listen(conn, workerResChan)

	// TODO: Block with sender fn, handling StartGame, etc
	select {}
}

func createInitConnectionRequest(payload ConnectRequestPayload) *generated.Request {
	initConnectionRequest := &generated.InitConnectionRequest{
		Id:       payload.Id,
		Username: payload.Username,
	}
	return &generated.Request{
		RequestType: generated.RequestType_INIT_CONNECTION,
		Request: &generated.Request_InitConnection{
			InitConnection: initConnectionRequest,
		},
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
