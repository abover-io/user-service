package user

import (
	"github.com/aboverio/user-service/validations"
	"github.com/google/uuid"
)

func EditPassword(userId *uuid.UUID, payload *validations.EditPasswordPayload) error {
	user, err := FindById(userId)
	if err != nil {
		return err
	}

	// @TODO: implement password change using bcrypt

	user.Password = payload.NewPassword

	return nil
}
