package privatechat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/test"
	"github.com/cucumber/godog"
	"github.com/gorilla/websocket"
)

type Message struct {
	Id      string
	Type    models.MessageType
	Target  string
	Content string
}

type Chat struct {
	sentMessages     map[string]any
	receivedMessages map[string]any
}

type privateChat struct {
	// t    testing.T
	resp httptest.ResponseRecorder
	test.TestInstance
	message Message
	users   *[]string
	chat    *Chat
	client  map[string]*websocket.Conn
}

func TestPrivateChat(t *testing.T) {
	pchat := privateChat{

		client: make(map[string]*websocket.Conn),
		chat: &Chat{
			sentMessages:     make(map[string]any),
			receivedMessages: make(map[string]any),
		},
		resp: *httptest.NewRecorder(),
	}

	t_instance := test.InitiateTest("../../")

	pchat.TestInstance = t_instance

	test_suite := godog.TestSuite{
		ScenarioInitializer: pchat.InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"feature/private_chat.feature"},
			TestingT: t,
		},
	}

	if test_suite.Run() != 0 {

		t.Fatal()

	}

}

func (p *privateChat) twoUsersAreOnline(tb *godog.Table) error {
	users := []string{}
	var msg Message

	for _, row := range tb.Rows[1:] {
		for _, cell := range row.Cells {
			users = append(users, cell.Value)
		}

	}
	p.users = &users

	server := httptest.NewServer(p.Sv)

	defer server.Close()

	wsURL := strings.Replace(server.URL, "http", "ws", 1)

	for _, client_id := range users {

		url := fmt.Sprintf("%s/v1/message?client_id=%s", wsURL, client_id)
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return fmt.Errorf("failed to connect WebSocket for user %s: %v", client_id, err)
		}
		p.client[client_id] = conn
		go func(clientID string) {
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					return
				}
				if len(message) == 0 {
					continue
				}
				err = json.Unmarshal(message, &msg)

				if err != nil {
					log.Println(err)
					return
				}

				p.chat.receivedMessages[clientID] = msg.Content
			}
		}(client_id)

	}

	return nil
}

func (p *privateChat) theFirstUserTextsAFriendWithAMessage(cht *godog.Table) error {

	var friend string
	var message string

	for _, row := range cht.Rows[1:] {
		friend = row.Cells[0].Value
		message = row.Cells[1].Value
		p.chat.sentMessages[friend] = message
		// expected message
		p.chat.receivedMessages[friend] = message

	}
	msg := Message{
		Id:      "msg_user1",
		Type:    models.PrivateMessage,
		Target:  friend,
		Content: message,
	}

	ws_msg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	log.Println("------------- TEXTING FRIEND--------")
	// sending message type
	return p.client[friend].WriteMessage(websocket.TextMessage, ws_msg)

}

func (p *privateChat) theFriendShouldReceiveTheSameMessage(cht *godog.Table) error {

	var message string
	var user_id string
	for _, row := range cht.Rows[1:] {
		user_id = row.Cells[0].Value
		message = row.Cells[1].Value

	}

	timeout := time.After(2 * time.Second)
	select {

	case <-timeout:
		return fmt.Errorf("client %s did not receive the expected message %q", user_id, p.chat.receivedMessages[user_id])

	default:
		if p.chat.receivedMessages[user_id] != message {
			return fmt.Errorf("message didn't match")

		}
	}
	r := p.chat.receivedMessages[user_id]
	log.Printf("recieved message::%v", r)

	return nil

}

func (p *privateChat) InitializeScenario(ctx *godog.ScenarioContext) {

	ctx.After(func(c context.Context, sc *godog.Scenario, err error) (context.Context, error) {

		for _, con := range p.client {
			con.Close()
		}

		return c, nil
	})

	ctx.Step(`^two users are online$`, p.twoUsersAreOnline)
	ctx.Step(`^the first user texts a friend with a message$`, p.theFirstUserTextsAFriendWithAMessage)
	ctx.Step(`^the friend should receive the same message$`, p.theFriendShouldReceiveTheSameMessage)
	ctx.Step(`^the friend texts the user back with a message$`, p.theFirstUserTextsAFriendWithAMessage)
	ctx.Step(`^the user receive the same message from friend$`, p.theFriendShouldReceiveTheSameMessage)
}
