package initiator

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/BoruTamena/go_chat/internal/constant/models/persistencedb"
	"github.com/gin-gonic/gin"
)

// @title Go Chat
// @version 1.0.0
// @description This is a  Swagger API documentation for GoChat Open source Project.
// @contact.name Boru Tamene Yadeta
// @contact.url  https://github.com/BoruTamena

func Init() {

	engine := gin.Default()

	rg := engine.Group("v1")

	// init config

	err, config := InitViper()

	if err != nil {
		log.Fatalf("Config::%v", err.Error())
	}

	// create logger
	logger := log.New(os.Stdout, "", 0)

	// init platform
	platform := InitPlatFormLayer()

	go platform.WebSocket.Run(context.Background())

	// db

	_, client := IntMgDb(*config)

	p_db := persistencedb.NewMgPersistence(&client, logger, *config)

	persistence := InitPersistence(p_db, *config)

	modules := InitModule(persistence.pchat, logger, platform)

	// init handler
	handler := IntHandler(logger, modules)

	InitRouter(*rg, handler, modules, platform)

	// listining and serving

	server := http.Server{
		Addr:    "localhost:" + config.Server.Port,
		Handler: engine,
	}

	logger.Fatal(server.ListenAndServe())

}
