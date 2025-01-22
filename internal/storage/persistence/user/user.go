package user

import (
	"context"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/db"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/BoruTamena/go_chat/internal/storage"
	"github.com/google/uuid"
)

type userStorage struct {
	db  persistencedb.MgPersistence
	cfg dto.Config
}

func NewUserStorage(mg persistencedb.MgPersistence, cfg dto.Config) storage.User {

	return &userStorage{
		db:  mg,
		cfg: cfg,
	}
}

func (u *userStorage) CreateUser(ctx context.Context, user dto.User) (db.User, error) {

	user_d, err := u.db.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		err = errors.WriteErr.Wrap(err, "insert new user err::").
			WithProperty(errors.ErrorCode, 500)

		log.Println("insert err::", err)

		return db.User{}, err
	}

	return user_d, nil

}

func (u *userStorage) GetUserByEmail(ctx context.Context, email string) (db.User, error) {

	user, err := u.db.GetUserByEmail(ctx, email)

	if err != nil {

		err = errors.DbReadErr.Wrap(err, "read user by email").
			WithProperty(errors.ErrorCode, 500)

		log.Println("Get user by email error::", err)
		return db.User{}, err
	}

	if user.ID.String() != "" {

		return user, nil

	}

	return db.User{}, errors.DbReadErr.New("no user with email")

}
