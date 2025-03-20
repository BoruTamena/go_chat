package friendship

import (
	"context"
	"fmt"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/internal/storage"
)

type friendshipModule struct {
	lg     *log.Logger
	user   storage.User
	friend storage.Friendship
}

func NewFriendShipModule(lg *log.Logger, userStorage storage.User,
	friendStorage storage.Friendship) module.Friendship {

	return &friendshipModule{
		lg:     lg,
		user:   userStorage,
		friend: friendStorage,
	}

}

// func (f *friendshipModule) GetFriendByUserName(ctx context.Context, user_name string) (dto.User, error) {

// }

func (f *friendshipModule) AddFriend(ctx context.Context, friend_user_name string) error {

	user := ctx.Value("user")

	user_dto, ok := user.(dto.User)

	if !ok {

		return fmt.Errorf("cant convert user from context to userdto object ")

	}
	// get friend
	friend, err := f.friend.GetFriendByUserName(ctx, friend_user_name)

	if err != nil {
		return err
	}

	// change
	err = f.friend.AddFriend(ctx, user_dto.Id, friend.Id)

	if err != nil {

		return err
	}

	return nil

}

func (f *friendshipModule) AcceptOrBlockFriend(ctx context.Context, friend_user_name, status string) error {

	user := ctx.Value("user")

	user_dto, ok := user.(dto.User)

	if !ok {

		return fmt.Errorf("cant convert user from context to userdto object ")

	}
	// get friend
	friend, err := f.friend.GetFriendByUserName(ctx, friend_user_name)

	if err != nil {
		return err
	}

	err = f.friend.UpdateFriendStatus(ctx, user_dto.Id, friend.Id, status)

	if err != nil {

		return err

	}

	return nil

}

func (f *friendshipModule) GetFriend(ctx context.Context, user_name string) (dto.User, error) {

	// get friend
	friend, err := f.friend.GetFriendByUserName(ctx, user_name)

	if err != nil {
		return dto.User{}, err
	}

	return friend, nil
}
