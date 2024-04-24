package routers

import (
	"time"

	"github.com/aboverio/user-service/routers/user"
	"github.com/gin-gonic/gin"
)

func Main(server *gin.Engine) {
	r := server.Group("/v1/users")

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success":   true,
			"timestamp": time.Now(),
		})
	})

	user.Router(r)

	server.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(404, gin.H{
			"success":   false,
			"timestamp": time.Now(),
			"data": gin.H{
				"message": "route not found",
			},
		})
	})
}
