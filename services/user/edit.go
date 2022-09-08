package user

import (
	"github.com/aboverio/user-service/services"
	"github.com/aboverio/user-service/validations"
	"github.com/google/uuid"
)

func Edit(userId *uuid.UUID, payload *validations.EditUserPayload) error {
	user, err := FindById(userId)
	if err != nil {
		return err
	}

	user.FirstName = payload.FirstName
	user.LastName = &payload.LastName

	if err := services.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}
