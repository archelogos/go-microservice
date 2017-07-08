package routes

import (
	"log"
	"net/http"

	"github.com/archelogos/go-microservice/microservice/services"
	"github.com/gin-gonic/gin"
)

// GetUserRouter get user
func GetUserRouter(c *gin.Context) {
	log.Println("[ROUTER] Getting user")
	name := c.Param("name")
	user := services.GetUser(name)
	c.JSON(http.StatusOK, user)
}
