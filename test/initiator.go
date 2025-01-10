package test

import (
	"context"
	"log"
	"os"

	"github.com/BoruTamena/go_chat/initiator"
	"github.com/gin-gonic/gin"
)

// your test initiator goes here

type TestInstance struct {
	// Server ,Cache
	// Module Handler
	// PlatformLayer
	Sv       *gin.Engine
	Handler  initiator.Handler
	Moudle   initiator.Module
	Platform initiator.PlatFormLayer
}

func InitiateTest(arg string) TestInstance {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)

	test_server := gin.Default()

	v1 := test_server.Group("v1")

	logger.Println("///>>> platform layer init...")
	platform := initiator.InitPlatFormLayer()

	logger.Println("///>> socket is listening to customer")
	// listen if the websocket client comes in
	go platform.WebSocket.Run(context.Background())

	logger.Println("///>>> module layer init...")
	module := initiator.InitModule(logger, platform)

	logger.Println("///>>> handler layer init...")
	handler := initiator.IntHandler(logger, module)

	logger.Println("///>>> route layer init...")
	initiator.InitRouter(*v1, handler, module, platform)

	logger.Println("///>>>  initilazation completed ")

	return TestInstance{
		Sv:       test_server,
		Handler:  handler,
		Moudle:   module,
		Platform: platform,
	}

}
