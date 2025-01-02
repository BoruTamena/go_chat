package initiator

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/internal/module/message"
)

type Module struct {

	/* all your modules goes here */

	MessageModule module.Message
}

func InitModule(lg *log.Logger, plt PlatFormLayer) Module {

	return Module{
		MessageModule: message.NewChatMessage(lg,
			plt.WebSocket),
	}

}
