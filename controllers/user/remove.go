package user

import (
	"net/http"
	"time"

	UserService "github.com/aboverio/user-service/services/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Remove(c *gin.Context) {
	userId, err := uuid.Parse(c.MustGet("user_id").(string))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := UserService.Remove(&userId); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"timestamp": time.Now(),
	})
}
