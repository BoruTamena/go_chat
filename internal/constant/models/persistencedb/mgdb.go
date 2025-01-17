package persistencedb

import (
	"log"

	"github.com/BoruTamena/go_chat/internal/constant/models/db"
	"github.com/BoruTamena/go_chat/internal/constant/models/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type MgPersistence struct {
	db.Queries
	client *mongo.Client
	lg     *log.Logger
	cfg    dto.Config
}

func NewMgPersistence(client *mongo.Client, lg *log.Logger, cfg dto.Config) MgPersistence {
	return MgPersistence{
		client: client,
		lg:     lg,
		cfg:    cfg,
	}
}
