package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/google/uuid"
)

func FindById(id *uuid.UUID) (*models.User, error) {
	user := &models.User{
		ID: *id,
	}

	if err := services.DB.First(user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
