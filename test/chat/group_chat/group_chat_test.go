package groupchat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/BoruTamena/go_chat/internal/constant/models"
	"github.com/BoruTamena/go_chat/test"
	"github.com/cucumber/godog"
	"github.com/gorilla/websocket"
)

type groupMembers struct {
	Id        string
	GroupName string
}
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

type groupChat struct {
	test.TestInstance
	chat    *Chat
	clients map[string]*websocket.Conn
	groups  []groupMembers
	mu      sync.Mutex
}

func TestGroupChat(t *testing.T) {

	t_instance := test.InitiateTest("../../")

	gc := groupChat{
		TestInstance: t_instance,
		clients:      make(map[string]*websocket.Conn),
		chat: &Chat{
			sentMessages:     make(map[string]any),
			receivedMessages: make(map[string]any),
		},
		groups: []groupMembers{},
	}

	// creating test suite
	test_suite := godog.TestSuite{
		ScenarioInitializer: gc.ScenarioInitializerfunc,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"feature/group_chat.feature"},
			TestingT: t,
		},
	}

	if test_suite.Run() != 0 {
		t.Fatal()
	}

}

func (gc *groupChat) allTheFollowingUsersAreActive(gm *godog.Table) error {

	for _, row := range gm.Rows[1:] {

		g := groupMembers{
			Id:        row.Cells[0].Value,
			GroupName: row.Cells[1].Value,
		}
		gc.groups = append(gc.groups, g)
		fmt.Println(gc.groups)
	}

	server := httptest.NewServer(gc.Sv)

	defer server.Close()

	WsUrl := strings.Replace(server.URL, "http", "ws", 1)

	group := make(map[string]bool)

	for _, client := range gc.groups {

		url := fmt.Sprintf("%s/v1/ws?client_id=%s", WsUrl, client.Id)
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return fmt.Errorf("failed to connect WebSocket for user %s: %v", client.Id, err)
		}
		gc.clients[client.Id] = conn

		if !group[client.GroupName] {

			err := gc.Platform.WebSocket.CreateRoom(context.Background(), client.Id, client.GroupName)

			if err != nil {
				return err
			}

			group[client.GroupName] = true

		}

		err = gc.Platform.WebSocket.JoinRoom(context.Background(), client.Id, client.GroupName)

		if err != nil {
			return err
		}

		go func(clientID string, conn *websocket.Conn) {
			var msg string
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					break

				}
				if len(message) == 0 {
					continue
				}
				err = json.Unmarshal(message, &msg)
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Println(msg)
				gc.mu.Lock()
				gc.chat.receivedMessages[clientID] = msg
				gc.mu.Unlock()
			}
		}(client.Id, conn)
	}

	return nil
}

func (gc *groupChat) theUserFromTheGroupSendTextMessage(cht *godog.Table) error {

	var friend string
	var message string
	var groupname string

	for _, row := range cht.Rows[1:] {
		friend = row.Cells[0].Value
		groupname = row.Cells[1].Value
		message = row.Cells[2].Value
		gc.chat.sentMessages[friend] = message

	}
	msg := Message{
		Id:      "msg_user_" + friend,
		Type:    models.GroupMessage,
		Target:  groupname,
		Content: message,
	}

	ws_msg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	log.Println("------------- TEXTING Group--------")
	fmt.Println(friend, message)
	// sending message type
	return gc.clients[friend].WriteMessage(websocket.TextMessage, ws_msg)

}

func (gc *groupChat) allMembersOfGroupShouldReceiveTheSameMessage(cht *godog.Table) error {

	time.Sleep(time.Second)
	var message string
	var group string
	for _, row := range cht.Rows[1:] {
		group = row.Cells[0].Value
		message = row.Cells[1].Value

	}

	for _, client := range gc.groups {

		if group != client.GroupName {
			continue
		}

		fmt.Println("--------recevied message -------")
		fmt.Println(client.Id)
		fmt.Println("===================")
		gc.mu.Lock()
		msg, ok := gc.chat.receivedMessages[client.Id]
		gc.mu.Unlock()
		if !ok {
			return fmt.Errorf("client %s did not receive any message", client.Id)
		}
		fmt.Println(msg)
		if msg != message {
			return fmt.Errorf("want:%v got:%v", message, msg)
		}

	}

	return nil
}

func (gc *groupChat) theUserFromDifferentGroupShouldNotReceiveTheMessage() error {

	other_group := gc.groups[3]
	gc.mu.Lock()
	msg, _ := gc.chat.receivedMessages[other_group.Id]
	gc.mu.Unlock()

	if msg != nil {
		return fmt.Errorf("Other group recieve hijack message :: %v", msg)
	}

	fmt.Println("---------------Other Group Member------------")
	fmt.Printf("msg:%v", msg)

	return nil
}

func (gc *groupChat) ScenarioInitializerfunc(sc *godog.ScenarioContext) {

	sc.After(func(c context.Context, sc *godog.Scenario, err error) (context.Context, error) {

		for _, con := range gc.clients {
			if con != nil {
				con.Close()
			}

		}

		return c, nil
	})

	sc.Step(`^all  members of  group  should receive the same message$`, gc.allMembersOfGroupShouldReceiveTheSameMessage)
	sc.Step(`^all the following users are active$`, gc.allTheFollowingUsersAreActive)
	sc.Step(`^the user from different group should not receive the message$`, gc.theUserFromDifferentGroupShouldNotReceiveTheMessage)
	sc.Step(`^the user from the group send text message$`, gc.theUserFromTheGroupSendTextMessage)

}
