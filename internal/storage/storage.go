package storage

import (
	"context"

	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
)

// define U storage interface here

type User interface {
}

type Chat interface {
	InsertChat(ctx context.Context, chat dto.Chat) error
	InsertGroupChat(ctx context.Context, gchat dto.GroupChat) error
}
