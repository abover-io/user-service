package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/aboverio/user-service/validations"
	"github.com/google/uuid"
)

func Edit(userId *uuid.UUID, payload *validations.EditUserPayload) error {
	user := &models.User{
		ID: *userId,
	}

	if err := services.DB.First(user).Error; err != nil {
		return errors.New("user not found")
	}

	user.FirstName = payload.FirstName
	user.LastName = &payload.LastName

	if err := services.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}
