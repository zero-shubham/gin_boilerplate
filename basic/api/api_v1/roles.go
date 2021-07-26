package apiv1

import (
	accesscontrol "basic/libs/access_control"

	"github.com/gin-gonic/gin"
)

func AddRoles(rg *gin.RouterGroup) {
	roles := rg.Group("/roles")

	roles.GET("/", getRolesHandler)

}

func getRolesHandler(c *gin.Context) {
	enfcr, err := accesscontrol.GetEnforcer()
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

// 	// c.JSON(200, gin)
// }
