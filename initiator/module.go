package initiator

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/module"
	"github.com/BoruTamena/go_chat/internal/module/message"
	"github.com/BoruTamena/go_chat/internal/module/user"
)

type Module struct {

	/* all your modules goes here */

	MessageModule module.Message

	UserModule module.User
}

func InitModule(stg Persistence, lg *log.Logger, plt PlatFormLayer) Module {

	return Module{
		MessageModule: message.NewChatMessage(lg, stg.Pchat, plt.WebSocket),
		UserModule:    user.NewUserManagement(lg, stg.User),
	}

}
