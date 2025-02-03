package login

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BoruTamena/go_chat/internal/helper"
	"github.com/BoruTamena/go_chat/seed"
	"github.com/BoruTamena/go_chat/test"
	"github.com/cucumber/godog"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userLogin struct {
	test.TestInstance
	user *User
	resp *http.Response
}

func TestUserLogin(t *testing.T) {

	t_instance := test.InitiateTest("./../../../")

	ur := userLogin{
		TestInstance: t_instance,
		resp:         &http.Response{},
	}

	// creating test suite
	test_suite := godog.TestSuite{
		ScenarioInitializer: ur.ScenarioInitializerfunc,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"feature/login_feature.feature"},
			TestingT: t,
		},
	}

	if test_suite.Run() != 0 {
		t.Fatal()
	}

}

func (ulg *userLogin) userIsARegisteredUser() error {

	seed.Excute(ulg.ConPool, "SeedUsers")

	return nil

}

func (ulg *userLogin) theUserLoginWithTheirCredential(arg *godog.Table) error {

	var user User

	for _, row := range arg.Rows[1:] {
		user.Email = row.Cells[0].Value
		user.Password = row.Cells[1].Value

	}

	ulg.user = &user

	fmt.Println("user login data----------->", ulg.user)
	fmt.Println(helper.HashPassword("password"))
	// defining server and sending request
	server := httptest.NewServer(ulg.Sv)

	defer server.Close()

	body, _ := json.Marshal(user)

	req, err := http.NewRequest(http.MethodPost, server.URL+"/v1/signin", bytes.NewReader(body))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	ulg.resp = res
	return nil
}

func (ulg *userLogin) theSystemShouldAuthorizeTheUser() error {

	fmt.Println("status code------------>", ulg.resp.StatusCode)

	if http.StatusCreated != ulg.resp.StatusCode {

		return errors.New("cant create this user ")
	}

	defer ulg.resp.Body.Close()

	b, _ := io.ReadAll(ulg.resp.Body)

	fmt.Println(string(b))

	return nil

}

func (ulg *userLogin) ScenarioInitializerfunc(sc *godog.ScenarioContext) {

	sc.After(func(c context.Context, sc *godog.Scenario, err error) (context.Context, error) {

		// resting response
		ulg.resp = &http.Response{}

		return c, nil
	})
	// steps
	sc.Step("^user is a registered user$", ulg.userIsARegisteredUser)
	sc.Step("^the user login with their credential$", ulg.theUserLoginWithTheirCredential)
	sc.Step("^the system should authorize the user$", ulg.theSystemShouldAuthorizeTheUser)

}
