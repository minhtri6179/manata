package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/minhtri6179/manata/component/tokenprovider"
)

type jwtProvider struct {
	serect string
}

func NewJwtProvider(serect string) *jwtProvider {
	return &jwtProvider{
		serect: serect,
	}
}

type TokenPayLoad struct {
	Uid   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayLoad) UserId() int {
	return p.Uid
}
func (p TokenPayLoad) Role() string {
	return p.URole
}

type myClaims struct {
	PayLoad TokenPayLoad `json:"payload"`
	jwt.StandardClaims
}
type token struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt int       `json:"expired_at"`
}

func (t *token) GetToken() string {
	return t.Token
}
func (j *jwtProvider) SecrectKey() string {
	return j.serect
}
func (j *jwtProvider) Generate(data tokenprovider.TokenPayLoad, expity int) (tokenprovider.Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		PayLoad: TokenPayLoad{
			Uid:   data.UserId(),
			URole: data.Role(),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(expity) * time.Minute).Unix(),
			IssuedAt:  now.Unix(),
			Id:        fmt.Sprintf("%d", now.UnixNano()),
		},
	})
	sognedtoken, err := t.SignedString([]byte(j.serect))
	if err != nil {
		return nil, err
	}
	return &token{
		Token:     sognedtoken,
		ExpiredAt: expity,
		CreatedAt: now,
	}, nil

}

func (j *jwtProvider) Validate(token string) (tokenprovider.TokenPayLoad, error) {
	t, err := jwt.ParseWithClaims(token, &myClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.serect), nil
	})
	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}
	if !t.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}
	claims, ok := t.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return claims.PayLoad, nil
}
