package apiv1

import (
	"basic/controllers"
	"basic/core/schemas"

	"github.com/gin-gonic/gin"
)

func AddUsers(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/", createUserHandler)

}

func createUserHandler(c *gin.Context) {
	var objIn schemas.CreateUser
	err := c.BindJSON(&objIn)
	if err != nil {
		c.JSON(
			422, gin.H{
				"message": "request body can't be processed",
			})
		return
	}

	user, controllerErr := controllers.CreateUser(&objIn, []string{"user"})
	if controllerErr != nil {
		c.JSON(
			int(controllerErr.Type), gin.H{
				"message": controllerErr.Meta,
			})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
