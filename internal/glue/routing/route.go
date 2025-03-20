package routing

import (
	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/platform"
	"github.com/gin-gonic/gin"
)

type Router struct {
	/*
		method string
		path string
		Handler gin.HandlerFunc
		Middlewares [] gin.HandlerFunc
		Permission

	*/
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

type SocketRouter struct {
	MsgType models.MessageType
	Handler platform.HandlerFunc
}

// associate handler with endpoints
func RegisterRoute(rgroup *gin.RouterGroup, routes []Router) {

	for _, route := range routes {

		var handler []gin.HandlerFunc

		handler = append(handler, route.Middlewares...)
		handler = append(handler, route.Handler)

		rgroup.Handle(route.Method, route.Path, handler...)

	}

}

// add socket handler to socket handler map
func RegisterSocketRoute(mn platform.WsManager, routes []SocketRouter) {

	for _, route := range routes {

		mn.AddHandler(string(route.MsgType), route.Handler)

	}

}
