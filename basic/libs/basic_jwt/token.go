package basicjwt

import (
	"basic/config"
	"basic/core/schemas"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func CreateToken(sub uuid.UUID, expiresIn int, issuer string) (string, error) {

	expAt := time.Now().Add(time.Minute * time.Duration(expiresIn))

	claims := schemas.TokenClaims{
		Sub: sub,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt.Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg := config.GetConfig()
	ss, err := token.SignedString([]byte(cfg.SigningKey))
	return ss, err
}

func DecodeToken(tokenString string) (*jwt.Token, error) {

	cfg := config.GetConfig()
	token, err := jwt.ParseWithClaims(tokenString, &schemas.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SigningKey), nil
	})

	return token, err
}
