package initiator

import (
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/BoruTamena/go_chat/internal/storage"
	"github.com/BoruTamena/go_chat/internal/storage/persistence/chat"
)

type Persistence struct {
	// privatechat
	Pchat storage.Chat
}

func InitPersistence(db persistencedb.MgPersistence, cfg dto.Config) Persistence {
	return Persistence{
		Pchat: chat.InitChat(db, cfg),
	}
}
