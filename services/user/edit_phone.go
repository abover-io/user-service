package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/aboverio/user-service/validations"
	"github.com/google/uuid"
)

func EditPhone(userId *uuid.UUID, payload *validations.EditPhonePayload) error {
	user := &models.User{
		ID: *userId,
	}

	if err := services.DB.First(user).Error; err != nil {
		return errors.New("user not found")
	}

	if *user.Phone == payload.Phone {
		return errors.New("new phone must not be the same as the old one")
	}

	user.Phone = &payload.Phone
	user.PhoneVerifiedAt = nil

	if err := services.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}
