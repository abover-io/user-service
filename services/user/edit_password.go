package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/aboverio/user-service/validations"
	"github.com/google/uuid"
)

func EditPassword(userId *uuid.UUID, payload *validations.EditPasswordPayload) error {
	user := &models.User{
		ID: *userId,
	}

	if err := services.DB.First(user).Error; err != nil {
		return errors.New("user not found")
	}

	// @TODO: implement password change using bcrypt

	return nil
}
