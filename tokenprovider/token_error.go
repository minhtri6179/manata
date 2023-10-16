package tokenprovider

import (
	"errors"

	"github.com/minhtri6179/manata/common"
)

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
