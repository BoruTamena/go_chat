package message

import (
	"context"
	"encoding/json"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/internal/storage"
	"github.com/BoruTamena/go_chat/platform"
)

type message struct {
	lg  *log.Logger
	Ws  platform.WsManager
	stg storage.Chat
}

func NewChatMessage(logger *log.Logger, messageStorage storage.Chat, ws platform.WsManager) module.Message {

	return &message{
		lg:  logger,
		Ws:  ws,
		stg: messageStorage,
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

	// TODO get sender id
	msg_dto := dto.Chat{
		SenderId:  "sender_id",
		ReciverId: message.Target,
		Message:   message.Content,
	}
	// storing data to the db

	if err := m.stg.InsertChat(ctx, msg_dto); err != nil {
		return err
	}
	return nil

}

func (m *message) MessageGroup(ctx context.Context, message models.Message) error {
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

	// TODO Check if client is online
	// If client is not online skip the broadcasting and store data
	// directly to database.
	if err := m.Ws.BroadCastMsgToRoom(ctx, message.Target, msg); err != nil {
		return err
	}

	// TODO Get senderId
	// storing data to database
	gchat := dto.GroupChat{
		SenderId:  "sender_id",
		GroupName: message.Target,
		Message:   message.Content,
	}
	if err := m.stg.InsertGroupChat(ctx, gchat); err != nil {
		return err
	}

	return nil
}
