package message

import (
	"net/http"

	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/internal/glue/routing"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/handler/middleware"
	"github.com/BoruTamena/go_chat/platform"
	"github.com/gin-gonic/gin"
)

func InitRoute(rg *gin.RouterGroup, mn platform.WsManager, mh handler.Message) {

	routes := []routing.Router{

		{
			Method:  http.MethodGet,
			Path:    "/ws",
			Handler: mn.ServeWs,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
		},

		{
			Method:  http.MethodGet,
			Path:    "/seemessage",
			Handler: mh.GetMessage,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
		},
	}

	routing.RegisterRoute(rg, routes)

}

func InitSocketRoute(mn platform.WsManager, handler handler.Message) {

	routes := []routing.SocketRouter{
		{
			MsgType: models.PrivateMessage,
			Handler: handler.TextFriendMessage,
		},
		{
			MsgType: models.GroupMessage,
			Handler: handler.TextGroupMessage,
		},
	}

	routing.RegisterSocketRoute(mn, routes)
}
