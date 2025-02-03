package middleware

import (
	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
)

func ErrorMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// handling error
		if len(c.Errors) > 0 {

			for _, err := range c.Errors {

				if err, ok := err.Err.(*errorx.Error); ok {

					code, _ := err.Property(errors.ErrorCode)

					StatusCode, _ := code.(int)

					c.AbortWithStatusJSON(StatusCode, gin.H{"err": err.Error()})

				}

			}
		}

	}

}
