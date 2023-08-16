package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xiaoxlm/miniprogram-demo/cmd/miniprogram/global"
	"time"
)

const (
	JwtKey = "rockontrol"
)

type CustomClaims struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	jwt.RegisteredClaims
}

func GenerateJWT() (string, error) {
	expire := time.Now().Add(7200 * time.Second) // 2小时

	var claims = CustomClaims{
		ClientID:     global.Config.ClientID,
		ClientSecret: global.Config.ClientSecret,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			Issuer:    "Xiaoxlm",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtKey))
}

func ParseJWT(tokenSTR string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenSTR, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("[CustomClaims] invalid")
}
