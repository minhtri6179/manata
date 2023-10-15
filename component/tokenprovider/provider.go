package tokenprovider

import (
	"errors"

	"github.com/minhtri6179/manata/common"
)

type Provider interface {
	// GetToken returns a token for the given service account.
	Generate(data TokenPayLoad, expity int) (Token, error)
	// ValidateToken validates the given token and returns the service account.
	Validate(token string) (TokenPayLoad, error)
	SecrectKey() string
}
type TokenPayLoad interface {
	UserId() int
	Role() string
}
type Token interface {
	GetToken() string
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"errNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"errEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token"),
		"invalid token",
		"errInvalidToken",
	)
	ErrTokenExpired = common.NewCustomError(
		errors.New("token expired"),
		"token expired",
		"errTokenExpired",
	)
	ErrInvalidTokenFormat = common.NewCustomError(
		errors.New("invalid token format"),
		"invalid token format",
		"errInvalidTokenFormat",
	)
)
