package message

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/platform"
	"github.com/gin-gonic/gin"
)

type message struct {
	lg        *log.Logger
	msgModule module.Message
}

func NewMessageHandler(lg *log.Logger, msg_m module.Message) handler.Message {

	return &message{
		lg:        lg,
		msgModule: msg_m,
	}
}

func (m *message) TextFriendMessage(ctx *gin.Context, message models.Message, client *platform.Client) {

	err := m.msgModule.MessageFriend(ctx, message)
	if err != nil {
		ctx.Error(err)
		return
	}
}
