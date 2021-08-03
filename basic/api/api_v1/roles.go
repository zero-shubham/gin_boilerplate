package apiv1

import (
	"basic/core/schemas"
	"basic/libs/middlewares"
	"basic/services"

	"github.com/gin-gonic/gin"
)

func AddRoles(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")

	roles.GET("/", middlewares.WithValidToken(getRolesHandler))

}

func getRolesHandler(c *gin.Context, token *schemas.TokenClaims) {
	enfcr, err := services.GetEnforcer()
	if err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Something went internally.",
			},
		)
		return
	}

	c.JSON(200, gin.H{
		"roles": enfcr.GetAllRoles(),
	})
}

// func addRoles(c *gin.Context) {
// 	enfcr, err := accesscontrol.GetEnforcer()
// 	if err != nil {
// 		c.JSON(
// 			500,
// 			gin.H{
// 				"message": "Something went internally.",
// 			},
// 		)
// 		return
// 	}

// 	enfcr.AddRoleForUser()
// }
