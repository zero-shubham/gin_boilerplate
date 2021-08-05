package apiv1

import (
	"basic/core/schemas"
	"basic/libs/middlewares"
	"basic/services"

	"github.com/gin-gonic/gin"
)

func AddRoles(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")

	roles.GET("/", middlewares.UserHasPermission(getRolesHandler, "roles", "read"))

}

func getRolesHandler(c *gin.Context, user *schemas.UserWithRoles) {
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
