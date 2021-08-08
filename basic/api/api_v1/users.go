package apiv1

import (
	"basic/controllers"
	"basic/core/schemas"
	"basic/libs/middlewares"
	"basic/libs/utils"
	"basic/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUsers(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/", createUserHandler)
	users.GET("/:id/roles", middlewares.WithAuthUser(getUserRolesHandler))

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

func getUserRolesHandler(c *gin.Context, user *schemas.UserWithRoles) {
	id := c.Param("id")

	if user.ID.String() != id && !utils.StringInSlice("root", user.Roles) {
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"message": "do not have enough permission for the action",
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

	roles, err := enfcr.GetRolesForUser(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Something went wrong internally.",
			},
		)
		return
	}

	c.JSON(200, gin.H{
		"roles": roles,
	})
}
