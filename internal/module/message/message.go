package message

import (
	"context"
	"encoding/json"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/platform"
)

type message struct {
	lg *log.Logger
	Ws platform.WsManager
}

func NewChatMessage(logger *log.Logger, ws platform.WsManager) module.Message {

	return &message{
		lg: logger,
		Ws: ws,
	}

}

func (m *message) MessageFriend(ctx context.Context, message models.Message) error {

	if err := message.Validate(); err != nil {

		err = errors.BadInput.Wrap(err, "invalid user input").
			WithProperty(errors.ErrorCode, 400)

		m.lg.Println("bad input:", err)
		return err
	}

	msg, err := json.Marshal(message.Content)

	if err != nil {

		err = errors.MarshalErr.Wrap(err, "can't marshal this message content").
			WithProperty(errors.ErrorCode, 500)

		m.lg.Println("unable to marshal message content", err)

		return err

	}

	if err := m.Ws.SendMessageToClient(ctx, message.Target, msg); err != nil {
		return err
	}

	return nil

}
