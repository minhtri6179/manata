package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/minhtri6179/manata/tokenprovider"
)

const minSecretKeySize = 32

type JWTProvider struct {
	secretKey string
}

func NewJWTProvider(secretKey string) (tokenprovider.TokenProvider, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secret key size: at least %d characters required", minSecretKeySize)
	}
	return &JWTProvider{secretKey}, nil
}

func (t *JWTProvider) GenerateToken(username string, expired time.Duration) (string, error) {
	payload, err := tokenprovider.NewPayLoad(username, expired)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(t.secretKey))
}
func (t *JWTProvider) ValidateToken(token string) (*tokenprovider.Payload, error) {
	key := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, tokenprovider.ErrInvalidToken
		}
		return []byte(t.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &tokenprovider.Payload{}, key)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && verr.Errors == jwt.ValidationErrorExpired {
			return nil, tokenprovider.ErrTokenExpired
		}
		return nil, tokenprovider.ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*tokenprovider.Payload)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	return payload, nil
}
