package ws

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/platform"

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

	if mn.rooms[room_name] == nil {
		mn.rooms[room_name] = make(map[string]*platform.Client)
	}

	mn.rooms[room_name][client_id] = client

	client.Rooms[room_name] = true

	return nil

}

func (mn *manager) JoinRoom(ctx context.Context, client_id, room_name string) error {

	mn.mu.Lock()

	defer mn.mu.Unlock()

	client, exists := mn.clients[client_id]

	if !exists {

		err := errors.CNotFound.New("client not found").WithProperty(errors.ErrorCode, 404)
		log.Println("client not found with id provided :", client_id)
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
	client.Rooms[room_name] = true
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

	mn.mu.Lock()
	room, exists := mn.rooms[room_name]
	mn.mu.Unlock()
	if !exists {
		err := errors.RoomErr.New("No Room Found ").WithProperty(errors.ErrorCode, 404)
		log.Printf("room not found :: %v", room_name)
		return err
	}

	if len(room) == 0 {
		log.Printf("room %s is empty, no clients to broadcast message", room_name)
		return nil
	}

	fmt.Println("-----------------Groups-----------")

	fmt.Println(room)

	var wg sync.WaitGroup
	for _, client := range room {

		wg.Add(1)

		go func(client *platform.Client) {

			defer wg.Done()

			select {
			case <-ctx.Done():
				log.Println("message broadcasting cancelled")
				return

			default:
				if err := client.Con.WriteMessage(websocket.TextMessage, message); err != nil {
					err = errors.ClientErr.Wrap(err, "failed to send message to client").WithProperty(errors.ErrorCode, 500)
					log.Println("failed to send message client", client)
				}

			}

		}(client)

	}

	wg.Wait()
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
