package util

import (
	"errors"

	"github.com/minhtri6179/manata/common"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash_password), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

var (
	ErrInvalidPassword = common.NewCustomError(
		errors.New("invalid password or username, please try again"),
		"invalid password error",
		"errInvalidPassword",
	)
)
