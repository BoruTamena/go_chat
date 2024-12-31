package ws

import (
	"context"
	"log"
	"sync"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/platform"
	"github.com/gorilla/websocket"
)

type manager struct {

	// map group names to the clients in that group
	// map of groups to client map
	groups map[string]map[string]*Client

	// track all the connected clients
	clients map[string]*Client

	mu sync.Mutex
}

func NewClientManger() platform.Manager {

	return &manager{
		groups:  make(map[string]map[string]*Client),
		clients: make(map[string]*Client),
		mu:      sync.Mutex{},
	}

}

func (mn *manager) AddClient(ctx context.Context, client_id string, conn *websocket.Conn) {

	client := &Client{
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
		err := errors.CNotFound.New("").WithProperty(errors.Key, 404)
		log.Print("No client with this id: ", err)
		return err
	}

	client.Con.Close()

	delete(mn.clients, client_id)

	mn.mu.Unlock()

	return nil

}
