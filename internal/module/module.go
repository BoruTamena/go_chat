package module

import (
	"context"

	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
)

type User interface {
	CreateUser(ctx context.Context, user dto.User) (dto.User, error)
	LogIn(ctx context.Context, user dto.UserLogin) (dto.User, error)
}

type Message interface {
	MessageFriend(ctx context.Context, message models.Message) error

	MessageGroup(ctx context.Context, message models.Message) error
}

type Friendship interface {
	// GetFriendByUserName(ctx context.Context, user_name string) (dto.User, error)
	AddFriend(ctx context.Context, friend_user_name string) error
	AcceptOrBlockFriend(ctx context.Context, friend_user_name, status string) error
}
