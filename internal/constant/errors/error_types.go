package errors

import "github.com/joomcode/errorx"

var (
	WsError = errorx.NewNamespace("WebSocket:error")
)

var (
	ClientErr = errorx.NewType(WsError, "Client Error")
	CNotFound = errorx.NewType(WsError, "Client Not Found ")
	RoomErr   = errorx.NewType(WsError, "Room Error")
)

var (
	ErrorCode = errorx.RegisterProperty("ERRCODE")
)
