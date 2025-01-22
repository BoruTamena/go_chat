package initiator

import (
	"context"
	"log"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"

	"github.com/jackc/pgx/v4/pgxpool"
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

func InitPgDb(cfg dto.Config) *pgxpool.Pool {

	url := cfg.Db.PgUrl

	con_config, err := pgxpool.ParseConfig(url)

	if err != nil {
		log.Fatalf("pg connection creation error:%v", err.Error())
	}

	con_config.ConnConfig.ConnectTimeout = time.Second

	// con_config.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {

	// 	return true
	// }
	// create connection pool

	con_pol, err := pgxpool.ConnectConfig(context.Background(), con_config)

	if err != nil {
		log.Fatalf("connection error::%v", err)
	}

	return con_pol

}
