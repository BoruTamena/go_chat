package friendship

import (
	"net/http"

	"github.com/BoruTamena/go_chat/internal/glue/routing"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func InitFriendShip(rg *gin.RouterGroup, handler handler.FriendShip) {

	route := []routing.Router{
		{

			Method:  http.MethodPost,
			Path:    "/friend",
			Handler: handler.GetFriendByUserName,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
		},
		{

			Method:  http.MethodPut,
			Path:    "/accept",
			Handler: handler.AcceptFriendRequest,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
		},
		{

			Method:  http.MethodPut,
			Path:    "/block",
			Handler: handler.BlockFriend,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
		},
	}

	routing.RegisterRoute(rg, route)

}
