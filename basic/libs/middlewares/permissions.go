package middlewares

import (
	"basic/core/models"
	"basic/core/schemas"
	basicjwt "basic/libs/basic_jwt"
	"basic/libs/utils"
	"basic/services"
	"net/http"
	"strings"
	"time"

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

func WithAuthUser(hf schemas.HandlerFuncWithAuthUser) gin.HandlerFunc {
	return WithValidToken(func(c *gin.Context, t *schemas.TokenClaims) {
		now := time.Now().Unix()

		if now > t.ExpiresAt {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Authorization token expired",
				},
			)
			return
		}

		// * get db
		db, err := services.GetDB()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "error getting DB conn",
				},
			)
			return
		}

		user, err := models.GetUserById(db, t.Sub)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Something went wrong internally.",
				},
			)
			return
		}
		if user == nil {
			c.JSON(
				http.StatusNotFound,
				gin.H{
					"message": "Invalid Authorization token",
				},
			)
			return
		}

		enfcr, err := services.GetEnforcer()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Something went wrong internally.",
				},
			)
			return
		}
		roles, err := enfcr.GetRolesForUser(t.Sub.String())
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Something went wrong internally.",
				},
			)
			return
		}

		hf(c, &schemas.UserWithRoles{
			Roles:     roles,
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	})
}

func UserHasPermission(hf schemas.HandlerFuncWithAuthUser, obj string, act string) gin.HandlerFunc {
	return WithAuthUser(func(c *gin.Context, u *schemas.UserWithRoles) {
		enfcr, err := services.GetEnforcer()
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Something went wrong internally.",
				},
			)
			return
		}

		ok, err := enfcr.Enforce(u.ID.String(), obj, act)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Something went wrong internally.",
				},
			)
			return
		}

		if !ok && !utils.StringInSlice("root", u.Roles) {
			c.JSON(
				http.StatusForbidden,
				gin.H{
					"message": "Not permitted.",
				},
			)
			return
		}

		hf(c, u)
	})
}
