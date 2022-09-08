package user

import (
	"net/http"
	"time"

	UserService "github.com/aboverio/user-service/services/user"
	"github.com/aboverio/user-service/validations"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EditEmail(c *gin.Context) {
	userId, err := uuid.Parse(c.MustGet("user_id").(string))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	payload := &validations.EditEmailPayload{}

	if err := c.BindJSON(payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := UserService.EditEmail(&userId, payload); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"timestamp": time.Now(),
	})
}
