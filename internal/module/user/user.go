package user

import (
	"context"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/helper"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/internal/storage"
)

type user struct {
	lg  *log.Logger
	stg storage.User
}

func NewUserManagement(lg *log.Logger, userStorage storage.User) module.User {

	return &user{
		lg:  lg,
		stg: userStorage,
	}

}

func (u *user) CreateUser(ctx context.Context, user dto.User) (dto.User, error) {

	if err := user.Validate(); err != nil {

		err = errors.BadInput.Wrap(err, "Invalid user input").
			WithProperty(errors.ErrorCode, 400)

		return dto.User{}, err

	}

	user_d, err := u.stg.GetUserByEmail(ctx, user.Email)

	if err == nil {

		err := errors.DublicateErr.New("user Exists").
			WithProperty(errors.ErrorCode, 409)

		log.Println("user exist err::", err)

		return dto.User{}, err

	}

	// if user_d.ID.String() != "" {

	// }

	h_password, err := helper.HashPassword(user.Password)

	if err != nil {
		return dto.User{}, err
	}

	user_dto := dto.User{
		UserName: user.UserName,
		Email:    user.Email,
		Password: h_password,
	}

	// call module storage here
	user_d, err = u.stg.CreateUser(ctx, user_dto)

	if err != nil {
		return dto.User{}, err
	}

	return dto.User{
		Id:       user_d.ID.String(),
		UserName: user_d.UserName,
		Email:    user_d.Email,
		Password: user_d.Password,
	}, nil

}

func (u *user) LogIn(ctx context.Context, user dto.UserLogin) (dto.User, error) {

	if err := user.Validate(); err != nil {
		err = errors.BadInput.Wrap(err, "Invalid user input").
			WithProperty(errors.ErrorCode, 400)

		return dto.User{}, err

	}

	// Get user
	user_d, err := u.stg.GetUserByEmail(ctx, user.Email)

	if err != nil {

		err = errors.AuthErr.Wrap(err, "Invalid email and password").
			WithProperty(errors.ErrorCode, 401)

		log.Println("login err::", err)

		return dto.User{}, err

	}

	if !helper.VerifyPassword(user.Password, user_d.Password) {
		err = errors.AuthErr.New("invalid email or password").
			WithProperty(errors.ErrorCode, 401)

		log.Println("login err::", err)
		return dto.User{}, err

	}

	return dto.User{
		Id:       user_d.ID.String(),
		UserName: user_d.UserName,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
