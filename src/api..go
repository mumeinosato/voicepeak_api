package src

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.POST("/tts", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			println(err.Error())
			return
		}

		wavBytes, err := Taks(task)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			println(err.Error())
			return
		}
		c.Data(200, "audio/wav", wavBytes)
	})
}