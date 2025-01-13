package persistencedb

import (
	"context"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupChat struct {
	ConverationId string    `bson:"conversation_id" json:"coversation_id"`
	GroupName     string    `bson:"group_name" json:"group_name"`
	SenderId      string    `bson:"sender_id" json:"sender_id"`
	Message       string    `bson:"message" json:"message"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}

func (g *MgPersistence) InsertGroupChat(ctx context.Context, gchat GroupChat) error {

	collection := g.client.Database(g.cfg.Db.Name).Collection("group_chat")

	res, err := collection.InsertOne(ctx, gchat)

	if err != nil {

		err = errors.WriteErr.Wrap(err, "inserting group chat error").
			WithProperty(errors.ErrorCode, 500)

		g.lg.Println("ERR::", err)

		return err

	}

	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {

		err = errors.NullObjId.New("").
			WithProperty(errors.ErrorCode, 500)

		g.lg.Println("ERR::", err)

		return err
	}

	return nil

}
