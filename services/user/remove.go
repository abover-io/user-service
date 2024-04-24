package user

import (
	"github.com/aboverio/user-service/services"
	"github.com/google/uuid"
)

func Remove(userId *uuid.UUID) error {
	user, err := FindById(userId)
	if err != nil {
		return err
	}

	if err := services.DB.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
