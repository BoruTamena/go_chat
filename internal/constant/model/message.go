package constant

type MessageType string

type Message struct {
	Id      string      `json:"id"`
	Type    MessageType `json:"type"`
	Target  string      `json:"target"`
	Content string      `json:"content"`
}
