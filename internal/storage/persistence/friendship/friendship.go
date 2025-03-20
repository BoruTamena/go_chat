package friendship

import (
	"context"
	"database/sql"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/db"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/BoruTamena/go_chat/internal/storage"
	"github.com/google/uuid"
)

type friendshipStorage struct {
	cfg dto.Config
	db  persistencedb.MgPersistence
}

func NewFriendShipStorage(cfg dto.Config, db persistencedb.MgPersistence) storage.Friendship {

	return &friendshipStorage{
		cfg: cfg,
		db:  db,
	}

}

func (fs *friendshipStorage) GetFriendByUserName(ctx context.Context, username string) (dto.User, error) {

	user, err := fs.db.GetFiendByUserName(ctx, username)

	if err != nil {

		err = errors.DbReadErr.Wrap(err, "no user with this username").
			WithProperty(errors.ErrorCode, 500)

		log.Printf("no user with a user_name :: %v", username)
		return dto.User{}, err
	}

	user_dto := dto.User{
		Id:       user.ID.String(),
		Email:    user.Email,
		Password: user.Password,
	}

	return user_dto, nil

}

func (fs *friendshipStorage) AddFriend(ctx context.Context, user_id, friend_id string) error {

	user_uuid, _ := uuid.Parse(user_id)
	friend_uuid, _ := uuid.Parse(friend_id)
	_, err := fs.db.AddFriend(ctx, db.AddFriendParams{
		UserID:   user_uuid,
		FriendID: friend_uuid,
		//TODO status
	})

	if err != nil {

		errors.WriteErr.Wrap(err, "cant add user to friendlist").
			WithProperty(errors.ErrorCode, 500)

		log.Printf("cant add this user to you friend list %v \n", friend_id)

		return err

	}

	return nil
}

func (fs *friendshipStorage) UpdateFriendStatus(ctx context.Context, user_id, friend_id, status string) error {

	user_uuid, _ := uuid.Parse(user_id)
	friend_uuid, _ := uuid.Parse(friend_id)
	_, err := fs.db.UpdateFriendStatus(ctx, db.UpdateFriendStatusParams{
		UserID:   user_uuid,
		FriendID: friend_uuid,
		Status:   sql.NullString{String: "accepted", Valid: true},
	})

	if err != nil {
		return err
	}

	return nil

}
