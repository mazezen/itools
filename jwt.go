package itools

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtClaims struct {
	Expire    time.Duration
	Secret    string
	LoginInfo interface{} `json:"login_info"`
	jwt.StandardClaims
}

func NewJwt(expire time.Duration, secret string) *JwtClaims {
	return &JwtClaims{
		Expire: expire,
		Secret: secret,
	}
}

// GenerateToken 签发Token
func (j *JwtClaims) GenerateToken(data interface{}) (string, error) {
	c := JwtClaims{
		LoginInfo: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.Expire).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(j.Secret))
}

// ParseToken 解析TOKEN
func (j *JwtClaims) ParseToken(token string) (*JwtClaims, error) {
	t, err := jwt.ParseWithClaims(token, &JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.Secret), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*JwtClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
