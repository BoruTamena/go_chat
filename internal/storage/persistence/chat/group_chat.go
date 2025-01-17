package chat

import (
	"context"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
)

func (g *chat) InsertGroupChat(ctx context.Context, gchat dto.GroupChat) error {

	c, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel()

	g_chat := persistencedb.GroupChat{

		ConverationId: "",
		GroupName:     gchat.GroupName,
		SenderId:      gchat.SenderId,
		Message:       gchat.Message,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := g.db.InsertGroupChat(c, g_chat); err != nil {
		return err
	}

	return nil

}
