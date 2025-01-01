package errors

import "github.com/joomcode/errorx"

var (
	InternalError = errorx.NewNamespace("Internal Server:error")
	InputError    = errorx.NewNamespace("InputValiation:error")
	WsError       = errorx.NewNamespace("WebSocket:error")
)

var (
	ClientErr    = errorx.NewType(WsError, "Client Error")
	CNotFound    = errorx.NewType(WsError, "Client Not Found ")
	RoomErr      = errorx.NewType(WsError, "Room Error")
	BadInput     = errorx.NewType(InputError, "Bad user input ")
	MarshalErr   = errorx.NewType(InternalError, "unable to marshal")
	UnMarshalErr = errorx.NewType(InternalError, "unable to unmarshal")
)

var (
	ErrorCode = errorx.RegisterProperty("ERRCODE")
)
