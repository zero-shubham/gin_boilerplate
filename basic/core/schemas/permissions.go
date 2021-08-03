package schemas

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type HandlerFuncWthToken func(c *gin.Context, t *TokenClaims)
type HandlerFuncWithAuthUser func(c *gin.Context, u *UserWithRoles)
type TokenClaims struct {
	Sub uuid.UUID `json:"sub"`
	jwt.StandardClaims
}
