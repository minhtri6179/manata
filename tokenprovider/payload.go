package tokenprovider

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	IssuedAt time.Time `json:"issued_at"`
	Expired  time.Time `json:"expired_at"`
}

func NewPayLoad(username string, expired time.Duration) (*Payload, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:       token,
		Username: username,
		IssuedAt: time.Now(),
		Expired:  time.Now().Add(expired),
	}
	return payload, nil

}
func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expired) {
		return ErrTokenExpired
	}
	return nil
}
