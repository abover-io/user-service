package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/google/uuid"
)

func Remove(userId *uuid.UUID) error {
	user := &models.User{
		ID: *userId,
	}

	if err := services.DB.First(user).Error; err != nil {
		return errors.New("user not found")
	}

	if err := services.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
