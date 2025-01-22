package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/gin-gonic/gin"
)

type user struct {
	lg         *log.Logger
	userModule module.User
}

func NewUserHandler(lg *log.Logger, user_module module.User) handler.User {

	return &user{
		lg:         lg,
		userModule: user_module,
	}
}

// @Summary create new user
// @Tags user
// @Description create new user on chat app
// @Accept json
// @Produce json
// @Param body body dto.User true "user request body"
// @Router /user [post]

func (u *user) RegisterUser(ctx *gin.Context) {

	fmt.Printf("============= user registeration is called ===============")

	var user dto.User

	if err := ctx.Bind(&user); err != nil {

		err = errors.BadInput.Wrap(err, "binding user input error").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return

	}

	// sending data to module
	n_user, err := u.userModule.CreateUser(ctx, user)

	if err != nil {
		ctx.Error(err)
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   n_user,
	})

}

// @Summary login to the system
// @Tags user
// @Description user login
// @Accept json
// @Produce json
// @Param body body dto.UserLogin true "user login request body"
// @Router /signin [post]

func (u *user) SignIn(ctx *gin.Context) {

	var user_lg dto.UserLogin

	if err := ctx.Bind(&user_lg); err != nil {

		err = errors.BadInput.Wrap(err, "binding user input error").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return

	}

	user, err := u.userModule.LogIn(ctx, user_lg)

	if err != nil {

		ctx.Error(err)
		return

	}

	// TODO create token and store it on

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})

}
