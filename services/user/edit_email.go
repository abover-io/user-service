package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/aboverio/user-service/validations"
	"github.com/google/uuid"
)

func EditEmail(userId *uuid.UUID, payload *validations.EditEmailPayload) error {
	user := &models.User{
		ID: *userId,
	}

	if err := services.DB.First(user).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Email == payload.Email {
		return errors.New("new email must not be the same as the old one")
	}

	user.Email = payload.Email
	user.EmailVerifiedAt = nil

	if err := services.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}
