package microservice

import "github.com/gin-gonic/gin"

// PingRouter ...
func PingRouter(c *gin.Context) {
	message := PingService("pong")
	c.JSON(200, gin.H{
		"message": message,
	})
}
