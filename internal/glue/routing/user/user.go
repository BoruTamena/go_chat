package user

import (
	"net/http"

	"github.com/BoruTamena/go_chat/internal/glue/routing"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/BoruTamena/go_chat/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(rg *gin.RouterGroup, handler handler.User) {

	route := []routing.Router{

		{
			Method:  http.MethodPost,
			Path:    "/signup",
			Handler: handler.RegisterUser,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
			},
		},

		{
			Method:  http.MethodPost,
			Path:    "/signin",
			Handler: handler.SignIn,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
			},
		},
	}

	routing.RegisterRoute(rg, route)

}
