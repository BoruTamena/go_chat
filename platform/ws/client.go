package ws

import (
	"context"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"

	"github.com/gorilla/websocket"
)

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

		err := errors.CNotFound.New("client not found").WithProperty(errors.ErrorCode, 404)
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

func (mn *manager) LeaveRoom(ctx context.Context, client_id, room_name string) error {

	mn.mu.Lock()

	defer mn.mu.Unlock()

	client, exists := mn.clients[client_id]

	if !exists {

		err := errors.CNotFound.New("client not found").WithProperty(errors.ErrorCode, 404)
		log.Println("client not found with id provided :", err)
		return err

	}

	room, ok := mn.rooms[room_name]

	if !ok {
		err := errors.RoomErr.New("No Room Found ").WithProperty(errors.ErrorCode, 404)
		log.Print("room not found", err)
		return err
	}
	// remove client from the room list
	delete(room, client_id)
	// remove room from client list
	delete(client.Rooms, room_name)

	return nil
}

func (mn *manager) BroadCastMsgToRoom(ctx context.Context, room_name string, message []byte) error {

	room, exists := mn.rooms[room_name]

	if !exists {
		err := errors.RoomErr.New("No Room Found ").WithProperty(errors.ErrorCode, 404)
		log.Print("room not found", err)
		return err
	}

	for _, client := range room {

		if err := client.Con.WriteMessage(websocket.TextMessage, message); err != nil {
			err = errors.ClientErr.Wrap(err, "failed to send message to client").WithProperty(errors.ErrorCode, 500)
			log.Println("failed to send message client", client)
			return err
		}

	}

	return nil
}

func (mn *manager) SendMessageToClient(ctx context.Context, client_id string, message []byte) error {

	client, ok := mn.clients[client_id]
	if !ok {
		err := errors.CNotFound.NewWithNoMessage().WithProperty(errors.ErrorCode, 404)
		log.Println("client not found with id provided :", err)
		return err
	}

	if err := client.Con.WriteMessage(websocket.TextMessage, message); err != nil {

		err = errors.ClientErr.Wrap(err, "could not write to client").WithProperty(errors.ErrorCode, 500)
		log.Println("could not write to this client", err)
		return err

	}

	return nil
}
