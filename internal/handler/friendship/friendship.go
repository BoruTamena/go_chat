package friendship

import (
	"log"
	"net/http"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/gin-gonic/gin"
)

type friendship struct {
	lg               *log.Logger
	friendshipModule module.Friendship
}

func NewFriendShipHandler(lg *log.Logger, f_module module.Friendship) handler.FriendShip {
	return &friendship{
		lg:               lg,
		friendshipModule: f_module,
	}
}

// @Summary getting friends all friendlists
// @Description get friends
// @Tags friend
// @Produce json
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Router /friends [get]
func (fh *friendship) GetFriends(ctx *gin.Context) {

}

// @Summary sending friend request
// @Description sending friend request to get connected with people
// @Tags friend
// @Accept json
// @Produce json
// @Param body body dto.FriendRequest true "friend request object"
// @Router /friend [post]
func (fh *friendship) GetFriendByUserName(ctx *gin.Context) {

	var req dto.FriendRequest

	if err := ctx.Bind(&req); err != nil {

		err = errors.BadInput.Wrap(err, "can't bind request").
			WithProperty(errors.ErrorCode, 500)

		ctx.Error(err)

		return
	}
	// adding user to friendlist
	err := fh.friendshipModule.AddFriend(ctx, req.FriendId)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": map[string]interface{}{
			"message": "friend request sent successfully",
		},
	})

}

// @Summary Accept friend request
// @Description accept request to stay connected with them regulary
// @Tags friend
// @Accept json
// @Produce json
// @Param body body dto.FriendUpdate true "friend update object"
// @Router /accept [put]
func (fh *friendship) AcceptFriendRequest(ctx *gin.Context) {

	if err := UpdateFriendStatus(ctx, fh); err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "user accepted",
	})

}

// @Summary Block friend request
// @Description block friend
// @Tags friend
// @Accept json
// @Produce json
// @Param body body dto.FriendUpdate true "friend update object"
// @Router /block [put]
func (fh *friendship) BlockFriend(ctx *gin.Context) {

	if err := UpdateFriendStatus(ctx, fh); err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "user blocked",
	})

}

func UpdateFriendStatus(ctx *gin.Context, fh *friendship) error {

	var friend dto.FriendUpdate

	if err := ctx.Bind(&friend); err != nil {

		err = errors.BadInput.Wrap(err, "failed to bind").
			WithProperty(errors.ErrorCode, http.StatusBadRequest)

		ctx.Error(err)

		return err

	}

	if err := fh.friendshipModule.
		AcceptOrBlockFriend(ctx, friend.UserName, friend.Status); err != nil {
		ctx.Error(err)
		return err
	}

	return nil

}
