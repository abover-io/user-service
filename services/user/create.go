package user

import (
	"errors"

	"github.com/aboverio/user-service/models"
	"github.com/aboverio/user-service/services"
	"github.com/aboverio/user-service/validations"
)

func Create(payload *validations.CreateUserPayload) (*models.User, error) {
	user := &models.User{
		Email: payload.Email,
	}

	if err := services.DB.First(user).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	user.FirstName = payload.FirstName
	user.LastName = &payload.LastName
	user.Phone = &payload.Phone
	user.Password = payload.Password

	if err := services.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
