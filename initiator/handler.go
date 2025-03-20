package initiator

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/handler/friendship"
	"github.com/BoruTamena/go_chat/internal/handler/message"
	"github.com/BoruTamena/go_chat/internal/handler/user"
)

type Handler struct {
	MessageHandler handler.Message
	UserHandler    handler.User

	FriendShipHandler handler.FriendShip
}

func IntHandler(lg *log.Logger, mdl Module) Handler {
	return Handler{
		MessageHandler:    message.NewMessageHandler(lg, mdl.MessageModule),
		UserHandler:       user.NewUserHandler(lg, mdl.UserModule),
		FriendShipHandler: friendship.NewFriendShipHandler(lg, mdl.friendshipModule),
	}
}
