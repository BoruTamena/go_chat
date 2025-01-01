package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type MessageType string

const (
	PrivateMessage MessageType = "PRIVATECHAT"
	GroupMessage   MessageType = "GROUPCHAT"
)

type Message struct {
	// message unique id
	Id string `json:"id"`
	// private / group message
	Type MessageType `json:"type"`
	// client_id / group_names
	Target string `json:"target"`
	// message contents
	Content string `json:"content"`
}

func (m Message) Validate() error {

	return validation.ValidateStruct(&m,
		validation.Field(&m.Id, validation.Required),
		validation.Field(&m.Type, validation.Required, validation.In(PrivateMessage, GroupMessage)),
		validation.Field(&m.Target, validation.Required),
		validation.Field(&m.Content, validation.Required, validation.Length(5, 1000)),
	)

}
