package middlewares

import (
	"basic/core/schemas"
	basicjwt "basic/libs/basic_jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithValidToken(hf schemas.HandlerFuncWthToken) gin.HandlerFunc {

	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		jwt := strings.Replace(bearerToken, "Bearer ", "", 1)
		token, err := basicjwt.DecodeToken(jwt)

		if err != nil || !token.Valid {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid Authorization token",
				},
			)
			return
		}
		claims := token.Claims.(*schemas.TokenClaims)
		hf(c, claims)
	}

}
