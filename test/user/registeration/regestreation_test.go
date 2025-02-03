package registeration

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BoruTamena/go_chat/test"
	"github.com/cucumber/godog"
)

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userRegistration struct {
	test.TestInstance
	user *User
	resp *http.Response
}

func TestUserRegistration(t *testing.T) {

	t_instance := test.InitiateTest("./../../../")

	ur := userRegistration{
		TestInstance: t_instance,
		resp:         &http.Response{},
	}

	// creating test suite
	test_suite := godog.TestSuite{
		ScenarioInitializer: ur.ScenarioInitializerfunc,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"feature/user_registration.feature"},
			TestingT: t,
		},
	}

	if test_suite.Run() != 0 {
		t.Fatal()
	}

}

func (ur *userRegistration) usersrequesttoregisterthemselfwithavaliddata(arg *godog.Table) error {

	var user User

	for _, row := range arg.Rows[1:] {

		user.UserName = row.Cells[0].Value
		user.Email = row.Cells[1].Value
		user.Password = row.Cells[2].Value

	}

	fmt.Println("read user=====> ", user)
	ur.user = &user

	// defining server and sending request
	server := httptest.NewServer(ur.Sv)

	defer server.Close()

	body, _ := json.Marshal(user)

	req, err := http.NewRequest(http.MethodPost, server.URL+"/v1/signup", bytes.NewReader(body))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	ur.resp = res
	return nil
}

func (ur *userRegistration) thesystemshouldregisterusersuccessfully() error {

	fmt.Println("status code------------>", ur.resp.StatusCode)

	if http.StatusCreated != ur.resp.StatusCode {

		return errors.New("cant create this user ")
	}

	return nil

}

func (ur *userRegistration) ScenarioInitializerfunc(sc *godog.ScenarioContext) {

	sc.After(func(c context.Context, sc *godog.Scenario, err error) (context.Context, error) {

		// resting response
		ur.resp = &http.Response{}

		return c, nil
	})

	// steps
	sc.Step("^users request to register themself with a valid data$", ur.usersrequesttoregisterthemselfwithavaliddata)
	sc.Step("^the system should register user successfully$", ur.thesystemshouldregisterusersuccessfully)

}
