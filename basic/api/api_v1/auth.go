package apiv1

import (
	"basic/controllers"
	"basic/core/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AddAuth(rg *gin.RouterGroup) {
	index := rg.Group("/auth")

	index.POST("/", loginHandler)

}

func loginHandler(c *gin.Context) {
	var loginCreds schemas.AuthPost
	err := c.ShouldBindWith(&loginCreds, binding.Form)
	if err != nil {
		c.JSON(
			http.StatusUnprocessableEntity, gin.H{
				"message": "request body can't be processed",
			})
		return
	}

	tokenResp, controllerErr := controllers.Login(&loginCreds)
	if controllerErr != nil {
		c.JSON(
			int(controllerErr.Type), gin.H{
				"message": controllerErr.Meta,
			})
		return
	}
	c.JSON(
		http.StatusOK, tokenResp,
	)
}
