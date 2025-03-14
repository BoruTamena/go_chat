package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// message handler
// @Tags message
// @Produce json
// @Router /seemessage [get]
func (m *message) GetMessage(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "reading message",
	})
}
