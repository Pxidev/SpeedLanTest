package routes

import "github.com/gin-gonic/gin"

func Setup(r *gin.Engine) {
	r.POST("/download", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
}
