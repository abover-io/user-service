package user

import (
	"net/http"
	"time"

	UserService "github.com/aboverio/user-service/services/user"
	"github.com/aboverio/user-service/validations"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	payload := &validations.CreateUserPayload{}

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := UserService.Create(payload)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":   true,
		"timestamp": time.Now(),
		"data":      user,
	})
}
