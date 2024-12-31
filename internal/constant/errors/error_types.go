package errors

import "github.com/joomcode/errorx"

var (
	WsError = errorx.NewNamespace("WebSocket:error")
)

var (
	CNotFound = errorx.NewType(WsError, "Client Not Found ")
)

var (
	Key = errorx.RegisterProperty("property_key")
)
