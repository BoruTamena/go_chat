package storage

import (
	"context"

	"github.com/BoruTamena/go_chat/internal/constant/models/db"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
)

// define U storage interface here

type User interface {
	CreateUser(ctx context.Context, user dto.User) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
}

type Chat interface {
	InsertChat(ctx context.Context, chat dto.Chat) error
	InsertGroupChat(ctx context.Context, gchat dto.GroupChat) error
}
