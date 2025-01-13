package errors

import "github.com/joomcode/errorx"

var (
	InternalError = errorx.NewNamespace("Internal Server:error")
	InputError    = errorx.NewNamespace("InputValiation:error")
	WsError       = errorx.NewNamespace("WebSocket:error")
	DbError       = errorx.NewNamespace("DataBase:error")
)

var (
	ClientErr = errorx.NewType(WsError, "Socket Client Error:error")
	CNotFound = errorx.NewType(WsError, "Socket Client Not Found:error ")

	WsConErr       = errorx.NewType(WsError, "Web SocketConnection:error")
	WsReadErr      = errorx.NewType(WsError, "Web Socket Read:error")
	WsUnRigsterErr = errorx.NewType(WsError, "Handler Not Regisetred:error")
	RoomErr        = errorx.NewType(WsError, "Room:error")
	BadInput       = errorx.NewType(InputError, "Bad user input:error")
	MarshalErr     = errorx.NewType(InternalError, "unable to marshal:error")
	UnMarshalErr   = errorx.NewType(InternalError, "unable to unmarshal:error")

	WriteErr  = errorx.NewType(DbError, "db write :: error ")
	NullObjId = errorx.NewType(DbError, "Null Object Id Returned :: error")
)

var (
	ErrorCode = errorx.RegisterProperty("ERRCODE")
)
