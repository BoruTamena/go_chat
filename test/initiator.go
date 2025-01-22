package test

import (
	"context"
	"log"
	"os"

	"github.com/BoruTamena/go_chat/initiator"
	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/gin-gonic/gin"
)

// your test initiator goes here

type TestInstance struct {
	// Server ,Cache
	// Module Handler
	// PlatformLayer
	Db       *persistencedb.MgPersistence
	Sv       *gin.Engine
	Handler  initiator.Handler
	Moudle   initiator.Module
	Platform initiator.PlatFormLayer
}

func InitiateTest(arg string) TestInstance {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)

	err, cfg := initiator.InitViper(arg)

	if err != nil {
		log.Fatal(err)
	}
	test_server := gin.Default()

	v1 := test_server.Group("v1")

	logger.Println("///>>> platform layer init...")
	platform := initiator.InitPlatFormLayer()

	logger.Println("///>> socket is listening to customer")
	// listen if the websocket client comes in
	go platform.WebSocket.Run(context.Background())

	_, db := initiator.IntMgDb(*cfg)

	con_pool := initiator.InitPgDb(*cfg)
	m := initiator.InitMigiration(arg+cfg.Migration.Path, cfg.Db.PgUrl)

	initiator.UpMigiration(m)

	p := persistencedb.NewMgPersistence(con_pool, &db, logger, *cfg)

	Stg := initiator.InitPersistence(p, *cfg)

	logger.Println("///>>> module layer init...")
	module := initiator.InitModule(Stg, logger, platform)

	logger.Println("///>>> handler layer init...")
	handler := initiator.IntHandler(logger, module)

	logger.Println("///>>> route layer init...")
	initiator.InitRouter(*v1, handler, module, platform)

	logger.Println("///>>>  initilazation completed ")

	return TestInstance{
		Db:       &p,
		Sv:       test_server,
		Handler:  handler,
		Moudle:   module,
		Platform: platform,
	}

}
