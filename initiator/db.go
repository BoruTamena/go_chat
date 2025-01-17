package initiator

import (
	"context"
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IntMgDb(cfg dto.Config) (error, mongo.Client) {

	url := cfg.Db.Url

	client_opt := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), client_opt)

	if err != nil {

		err := errors.DbError.NewType("mongodb Client:err").WrapWithNoMessage(err).
			WithProperty(errors.ErrorCode, 500)

		log.Println("unable to create a mongo db client::", err)
		return err, mongo.Client{}

	}
	// check connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		return err, mongo.Client{}
	}

	return nil, *client

}
