package initiator

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/handler/message"
)

type Handler struct {
	MessageHandler handler.Message
}

func IntHandler(lg *log.Logger, mdl Module) Handler {
	return Handler{
		MessageHandler: message.NewMessageHandler(lg, mdl.MessageModule),
	}
}
