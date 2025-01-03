package test

import (
	"log"
	"os"

	"github.com/BoruTamena/go_chat/initiator"
	"github.com/gin-gonic/gin"
)

// your test initiator goes here

type TestInstance struct {
	/*

		Server
		Cache
		Module
		Handler
		PlatformLayer

	*/
	Sv       *gin.Engine
	Handler  initiator.Handler
	Moudle   initiator.Module
	Platform initiator.PlatFormLayer
}

func InitiateTest(arg string) TestInstance {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)

	test_server := gin.Default()

	v1 := test_server.Group("v1")

	platform := initiator.InitPlatFormLayer()

	module := initiator.InitModule(logger, platform)

	handler := initiator.IntHandler(logger, module)

	initiator.InitRouter(*v1, handler, module, platform)

	return TestInstance{
		Sv:       test_server,
		Handler:  handler,
		Moudle:   module,
		Platform: platform,
	}

}
