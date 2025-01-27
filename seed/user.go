package seed

import (
	"fmt"

	"github.com/BoruTamena/go_chat/internal/helper"
)

func (s *Seed) SeedUsers() {

	pass, _ := helper.HashPassword("password")

	fmt.Println("hash password =", pass)

	Data := []SeedData{
		{
			TableName: "users",
			Columns:   []string{"user_name", "email", "password"},
			Values:    []interface{}{"user1", "user@gmail.com", pass},
		},
	}

	for _, item := range Data {

		s.SeedTable(item.TableName, item.Columns, item.Values)

	}

}
