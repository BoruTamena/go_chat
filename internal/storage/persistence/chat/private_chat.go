package chat

import (
	"context"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/BoruTamena/go_chat/internal/storage"
)

type chat struct {
	db  persistencedb.MgPersistence
	cfg dto.Config
}

func InitChat(mg persistencedb.MgPersistence, cfg dto.Config) storage.Chat {

	return &chat{
		db:  mg,
		cfg: cfg,
	}

}
func (p *chat) InsertChat(ctx context.Context, chat dto.Chat) error {

	c, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	p_chat := persistencedb.PrivateChat{
		ConversationId: "",
		SenderId:       chat.SenderId,
		ReciverId:      chat.ReciverId,
		Message:        chat.Message,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := p.db.InsertChat(c, p_chat); err != nil {
		return err
	}

	return nil

}
