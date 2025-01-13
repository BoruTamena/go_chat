package persistencedb

import (
	"context"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PrivateChat struct {
	ConversationId string    `bson:"conversation_id" json:"conversation_id,required"`
	SenderId       string    `bson:"sender_id" json:"sender_id"`
	ReciverId      string    `bson:"receiver_id" json:"receiver_id"`
	Message        string    `bson:"message" json:"message "`
	CreatedAt      time.Time `bson:"created_at" json:"chat_at"`
	UpdatedAt      time.Time `bson:"updated_at" json:"updated_at"`
}

func (p *MgPersistence) InsertChat(ctx context.Context, chat PrivateChat) error {

	collection := p.client.Database(p.cfg.Db.Name).Collection("private_chat")

	res, err := collection.InsertOne(ctx, chat)

	if err != nil {

		err = errors.WriteErr.Wrap(err, "inserting private chat error").
			WithProperty(errors.ErrorCode, 500)
		p.lg.Printf("ERR::", err)
		return err

	}

	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		err = errors.NullObjId.New("").
			WithProperty(errors.ErrorCode, 500)

		p.lg.Println("ERR::", err)

		return err
	}

	return nil

}
