package persistencedb

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/models/db"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgPersistence struct {
	*db.Queries
	client *mongo.Client
	lg     *log.Logger
	cfg    dto.Config
}

func NewMgPersistence(pool *pgxpool.Pool, client *mongo.Client, lg *log.Logger, cfg dto.Config) MgPersistence {
	return MgPersistence{
		Queries: db.New(pool),
		client:  client,
		lg:      lg,
		cfg:     cfg,
	}
}
