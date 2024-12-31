package errors

import "github.com/joomcode/errorx"

var (
	WsError = errorx.NewNamespace("WebSocket:error")
)

var (
	CNotFound = errorx.NewType(WsError, "Client Not Found ")
	RoomErr   = errorx.NewType(WsError, "Room Error")
)

var (
	ErrorCode = errorx.RegisterProperty("ERRCODE")
)
