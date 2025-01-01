package module

import (
	"context"

	"github.com/BoruTamena/go_chat/internal/constant/models"
)

type Message interface {
	MessageFriend(ctx context.Context, message models.Message) error
}
