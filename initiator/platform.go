package initiator

import (
	"github.com/BoruTamena/go_chat/platform"
	"github.com/BoruTamena/go_chat/platform/ws"
)

type PlatFormLayer struct {

	/*

	 sms platform.SMS

	 sso platfrom.SSO

	*/

	WebSocket platform.WsManager
}

func InitPlatFormLayer() PlatFormLayer {

	return PlatFormLayer{
		WebSocket: ws.NewClientManger(),
	}
}
