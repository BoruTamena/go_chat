package platform

import (
	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HandlerFunc func(ctx *gin.Context, Message models.Message, client *Client)

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
