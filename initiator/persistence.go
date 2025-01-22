package initiator

import (
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/BoruTamena/go_chat/internal/storage"
	"github.com/BoruTamena/go_chat/internal/storage/persistence/chat"
	"github.com/BoruTamena/go_chat/internal/storage/persistence/user"
)

type Persistence struct {
	// privatechat
	Pchat storage.Chat

	User storage.User
}

func InitPersistence(db persistencedb.MgPersistence, cfg dto.Config) Persistence {
	return Persistence{
		Pchat: chat.InitChat(db, cfg),
		User:  user.NewUserStorage(db, cfg),
	}
}
