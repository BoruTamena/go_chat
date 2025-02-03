package initiator

import (
	"github.com/BoruTamena/go_chat/docs"
	"github.com/BoruTamena/go_chat/internal/glue/routing/friendship"
	"github.com/BoruTamena/go_chat/internal/glue/routing/message"
	"github.com/BoruTamena/go_chat/internal/glue/routing/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(

	groupRouter gin.RouterGroup,
	handler Handler,
	module Module,
	platform PlatFormLayer) {

	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/v1"
	groupRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.InitRoute(&groupRouter, handler.UserHandler)

	// initalizing route
	message.InitRoute(&groupRouter, platform.WebSocket, handler.MessageHandler)
	message.InitSocketRoute(platform.WebSocket, handler.MessageHandler)

	// initalizing friendship route

	friendship.InitFriendShip(&groupRouter, handler.FriendShipHandler)

}
