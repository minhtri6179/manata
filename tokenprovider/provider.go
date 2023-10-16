package tokenprovider

import "time"

type TokenProvider interface {
	GenerateToken(username string, expired time.Duration) (string, error)
	ValidateToken(token string) (*Payload, error)
}
