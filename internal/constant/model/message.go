package constant

type MessageType string

const (
	PrivateMessage MessageType = "PRIVATECHAT"
	GroupMessage   MessageType = "GROUPCHAT"
)

type Message struct {
	Id      string      `json:"id"`
	Type    MessageType `json:"type"`
	Target  string      `json:"target"`
	Content string      `json:"content"`
}
