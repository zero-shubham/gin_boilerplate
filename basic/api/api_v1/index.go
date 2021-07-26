package apiv1

import (
	"github.com/gin-gonic/gin"
)

func AddIndex(rg *gin.RouterGroup) {
	index := rg.Group("/index")

	index.GET("/", getIndexHandler)

}

func getIndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello World!",
	})
}
