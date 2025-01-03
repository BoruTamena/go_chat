package initiator

import (
	"log"
	"net/http"
	"os"

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

	// create logger
	logger := log.New(os.Stdout, "", 0)

	//
	platform := InitPlatFormLayer()

	modules := InitModule(logger, platform)

	// init handler
	handler := IntHandler(logger, modules)

	InitRouter(*rg, handler, modules, platform)

	// listining and serving

	server := http.Server{
		Addr:    "localhost:7000",
		Handler: engine,
	}

	logger.Fatal(server.ListenAndServe())

}
