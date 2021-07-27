package basicjwt

import (
	"basic/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func CreateToken(sub uuid.UUID, expiresIn int, issuer string) (string, error) {
	type TokenClaims struct {
		Sub uuid.UUID `json:"sub"`
		jwt.StandardClaims
	}
	expAt := time.Now().Add(time.Minute * time.Duration(expiresIn))

	claims := TokenClaims{
		sub,
		jwt.StandardClaims{
			ExpiresAt: expAt.Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg := config.GetConfig()
	ss, err := token.SignedString([]byte(cfg.SigningKey))
	return ss, err
}
