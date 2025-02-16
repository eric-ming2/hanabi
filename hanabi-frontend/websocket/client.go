package websocket

import (
	"log"
	"net/url"

	"github.com/eric-ming2/hanabi/hanabi-frontend/generated"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type WorkerRequestType int

const (
	ConnectRequest WorkerRequestType = iota
	ReadyRequest
	StartGameRequest
)

type WorkerRequest struct {
	Type    WorkerRequestType
	Payload interface{}
}

type ConnectRequestPayload struct {
	Id       string
	Username string
}

type WorkerResponseType int

const (
	ConnectFailed WorkerResponseType = iota
	UpdateGameState
)

type WorkerResponse struct {
	Type    WorkerResponseType
	Payload interface{}
}

func ClientWorker(workerReqChan chan WorkerRequest, workerResChan chan WorkerResponse, id string) {
	var conn *websocket.Conn = nil
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	for req := range workerReqChan {
		switch req.Type {
		case ConnectRequest:
			payload, ok := req.Payload.(ConnectRequestPayload)
			if !ok {
				log.Fatalf("Failed to cast payload to ConnectRequestPayload")
			}
			conn = connect(workerResChan, payload)
		case ReadyRequest:
			if conn == nil {
				log.Fatalf("Tried to ready before conn initialized. This should be impossible.")
			}
			startGameReq, err := proto.Marshal(createReadyRequest(id))
			if err != nil {
				log.Fatalf("Failed to marshal proto: %s", err)
			}
			err = conn.WriteMessage(websocket.BinaryMessage, startGameReq)
			if err != nil {
				log.Fatalf("Failed to write ready message: %s", err)
			}
		case StartGameRequest:
			if conn == nil {
				log.Fatalf("Tried to start game before conn initialized. This should be impossible.")
			}
			startGameReq, err := proto.Marshal(createStartGameRequest(id))
			if err != nil {
				log.Fatalf("Failed to marshal proto: %s", err)
			}
			err = conn.WriteMessage(websocket.BinaryMessage, startGameReq)
			if err != nil {
				log.Fatalf("Failed to write start game message: %s", err)
			}
		}
	}
}

func connect(workerResChan chan WorkerResponse, payload ConnectRequestPayload) *websocket.Conn {
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
	log.Println("Connected to WebSocket server")

	initConnectionReq, err := proto.Marshal(createInitConnectionRequest(payload))
	if err != nil {
		log.Fatalf("Failed to marshal proto", err)
	}
	err = conn.WriteMessage(websocket.BinaryMessage, initConnectionReq)
	if err != nil {
		log.Fatalf("Failed to write init connection message", err)
	}

	//startGameReq, err := proto.Marshal(createStartGameRequest())
	//if err != nil {
	//	log.Fatalf("Failed to marshal proto", err)
	//}

	// TODO: Move into sender fn, send on button click
	//err = conn.WriteMessage(websocket.BinaryMessage, startGameReq)
	//if err != nil {
	//	log.Fatalf("Failed to write start game message", err)
	//}

	go listen(conn, workerResChan)

	return conn
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

func createReadyRequest(id string) *generated.Request {
	return &generated.Request{
		Id:          id,
		RequestType: generated.RequestType_READY,
		Request: &generated.Request_Ready{
			Ready: &generated.ReadyRequest{},
		},
	}
}

func createStartGameRequest(id string) *generated.Request {
	return &generated.Request{
		Id:          id,
		RequestType: generated.RequestType_START_GAME,
		Request: &generated.Request_StartGame{
			StartGame: &generated.StartGameRequest{},
		},
	}
}
