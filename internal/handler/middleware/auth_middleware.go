package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var user dto.User

		bears := "Bearer "

		auth_token := ctx.GetHeader("Authorization")

		if auth_token == "" || !strings.HasPrefix(auth_token, bears) {

			err := errors.AuthErr.New("unauth").
				WithProperty(errors.ErrorCode, 401)
			log.Println("auth err ::", err)
			ctx.Error(err)
			ctx.Abort()
			return

		}

		token := auth_token[len(bears):]

		p_token, _ := helper.ParseAccessToken(token)

		claim, err := p_token.Claims.GetSubject()

		if err != nil {
			//error

			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
			return
		}

		if err := json.Unmarshal([]byte(claim), &user); err != nil {
			// err
			ctx.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
			return
		}

		ctx.Set("user", user)

		// allow
		ctx.Next()

	}
}
