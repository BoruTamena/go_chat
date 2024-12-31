package ws

import (
	"context"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/gorilla/websocket"
)

type Client struct {
	// unique to each client
	ClientId string `json:"client_id"`
	// gorilla socket connection
	Con *websocket.Conn
	// meta data about client
	// status : online ,typing,and other
	MetaData map[string]interface{}
	// rooms the client has joined
	Rooms map[string]bool
}

func (mn *manager) CreateRoom(ctx context.Context, client_id, room_name string) error {

	mn.mu.Lock()
	defer mn.mu.Unlock()

	client, ok := mn.clients[client_id]

	if !ok {
		err := errors.CNotFound.NewWithNoMessage().WithProperty(errors.ErrorCode, 404)
		log.Println("client not found with id provided :", err)
		return err
	}

	// check if room exists with a given name

	if _, exists := mn.rooms[room_name]; exists {

		err := errors.RoomErr.New("room already exists").WithProperty(errors.ErrorCode, 409)
		log.Println("room exist with this name", err)
		return err

	}

	mn.rooms[room_name][client_id] = client

	return nil

}

func (mn *manager) JoinRoom(ctx context.Context, client_id, room_name string) error {

	mn.mu.Lock()

	defer mn.mu.Unlock()

	client, exists := mn.clients[client_id]

	if !exists {

		err := errors.CNotFound.NewWithNoMessage().WithProperty(errors.ErrorCode, 404)
		log.Println("client not found with id provided :", err)
		return err

	}

	// if room exists
	_, exists = mn.rooms[room_name]

	if !exists {
		err := errors.RoomErr.New("No Room Found ").WithProperty(errors.ErrorCode, 404)
		log.Print("room not found", err)
		return err
	}

	// add client to the room
	mn.rooms[room_name][client_id] = client

	return nil

}
