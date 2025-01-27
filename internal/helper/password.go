package helper

import (
	"strings"

	"github.com/BoruTamena/go_chat/internal/constant/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)),
		bcrypt.DefaultCost)

	if err != nil {

		err = errors.InternalError.NewType("password encription::err").
			Wrap(err, "can't hash password").
			WithProperty(errors.ErrorCode, 500)

		return "", err

	}
	return string(bytes), err

}

func VerifyPassword(password, hash_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(strings.TrimSpace(hash_password)),
		[]byte(strings.TrimSpace(password)))
	return err == nil

}
