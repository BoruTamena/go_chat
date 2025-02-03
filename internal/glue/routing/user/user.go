package user

import (
	"net/http"

	"github.com/BoruTamena/go_chat/internal/glue/routing"
	"github.com/BoruTamena/go_chat/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRoute(rg *gin.RouterGroup, handler handler.User) {

	route := []routing.Router{

		{
			Method:  http.MethodPost,
			Path:    "/signup",
			Handler: handler.RegisterUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/signin",
			Handler: handler.SignIn,
		},
	}

	routing.RegisterRoute(rg, route)

}
