package platform

import (
	"context"

	"github.com/gorilla/websocket"
)

// define your platform interfaces here

type SMS interface {
	// SMS Interface
}

type SSO interface {

	// SSO interface
}

type Manager interface {
	// web socket interface

	AddClient(ctx context.Context, client_id string, conn *websocket.Conn)
}
