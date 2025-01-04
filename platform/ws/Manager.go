package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models"

	"github.com/BoruTamena/go_chat/platform"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// ToDo CheckOrigin is used to verify
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// type HandlerFunc func(ctx context.Context, Message models.Message, client *Client)
type manager struct {

	// maps room names to the clients in that room
	// map of rooms to client map
	rooms map[string]map[string]*platform.Client
	// track all the connected clients
	clients map[string]*platform.Client
	// maps message handler type to handler
	handlers map[string]platform.HandlerFunc
	// register client
	register chan *platform.Client
	// un register client
	unregister chan *platform.Client

	mu sync.Mutex
}

func NewClientManger() platform.WsManager {

	return &manager{
		rooms:      make(map[string]map[string]*platform.Client),
		clients:    make(map[string]*platform.Client),
		register:   make(chan *platform.Client),
		unregister: make(chan *platform.Client),
		handlers:   make(map[string]platform.HandlerFunc),
		// mu:         sync.Mutex{},

	}

}

func (mn *manager) Run(ctx context.Context) {

	select {
	case client := <-mn.register:
		mn.AddClient(ctx, client.ClientId, client.Con)

	case client := <-mn.unregister:
		mn.RemoveClient(ctx, client.ClientId)
	}

}

func (mn *manager) AddClient(ctx context.Context, client_id string, conn *websocket.Conn) {

	client := &platform.Client{
		ClientId: client_id,
		Con:      conn,
		MetaData: make(map[string]interface{}),
		Rooms:    make(map[string]bool),
	}

	mn.mu.Lock()

	mn.clients[client_id] = client

	mn.mu.Unlock()

}

func (mn *manager) RemoveClient(ctx context.Context, client_id string) error {

	mn.mu.Lock()

	client, exists := mn.clients[client_id]

	if !exists {
		err := errors.CNotFound.New("").WithProperty(errors.ErrorCode, 404)
		log.Print("No client with this id: ", err)
		return err
	}

	client.Con.Close()

	for room_name := range client.Rooms {

		delete(mn.rooms[room_name], client_id)
		// delete the whole room if their no client left
		if len(mn.rooms[room_name]) == 0 {
			delete(mn.rooms, room_name)
		}
	}

	delete(mn.clients, client_id)

	mn.mu.Unlock()

	return nil

}

// this method meant to be
// imported directly from the ws package
func (mn *manager) AddHandler(chat_type string, handler platform.HandlerFunc) {

	mn.handlers[chat_type] = handler

}

func (mn *manager) ServeWs(ctx *gin.Context) {

	var message models.Message

	client_id := ctx.Query("client_id")
	if client_id == "" {
		err := errors.BadInput.New("client_id is required").
			WithProperty(errors.ErrorCode, 400)
		ctx.AbortWithError(400, err)
		return
	}

	webcon, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		err = errors.WsConErr.Wrap(err, "can't create connection")
		log.Println("web socket connection error", err)
		return
	}

	client := &platform.Client{
		// todo client_id should be unique
		ClientId: client_id,
		Con:      webcon,
	}

	mn.register <- client // register  client

	for {

		_, payload, err := webcon.ReadMessage()

		if err != nil {
			err = errors.WsReadErr.Wrap(err, "unable to read message from socket").
				WithProperty(errors.ErrorCode, 500)
			log.Println("Read error", err)
			break
		}

		if len(payload) == 0 {
			// no message sent
			continue
		}

		err = json.Unmarshal(payload, &message)
		if err != nil {
			err = errors.UnMarshalErr.Wrap(err, "unable to unmarshal payload to message").
				WithProperty(errors.ErrorCode, 500)

			log.Println("unmarshal error triggered", err)
			break
		}

		handler, ok := mn.handlers[string(message.Type)]
		if !ok {
			err := errors.WsUnRigsterErr.New("message handler not exist ").
				WithProperty(errors.ErrorCode, 500)
			log.Println(err)
			break
		}

		cl, ok := mn.clients[message.Target]
		if !ok {

			err := errors.CNotFound.New("no cliend with this client_id").
				WithProperty(errors.ErrorCode, 400)

			log.Printf(`client_id: %v`, message.Target, err)
			break

		}

		handler(ctx, message, cl)

	}

}
