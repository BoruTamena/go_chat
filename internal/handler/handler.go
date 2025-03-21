package handler

import (
	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/platform"
	"github.com/gin-gonic/gin"
)

// define your handlers interface here

type User interface {
	RegisterUser(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}
type Message interface {
	GetMessage(ctx *gin.Context)
	TextFriendMessage(ctx *gin.Context, message models.Message, client *platform.Client)
	TextGroupMessage(ctx *gin.Context, message models.Message, _ *platform.Client)
}

type FriendShip interface {
	GetFriends(ctx *gin.Context)
	GetFriendByUserName(ctx *gin.Context)
	AcceptFriendRequest(ctx *gin.Context)
	BlockFriend(ctx *gin.Context)
}
