package validations

type CreateUserPayload struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Password  string `json:"password" binding:"required"`
}

type EditUserPayload struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
}

type EditEmailPayload struct {
	Email string `json:"email" binding:"required,email"`
}

type EditPhonePayload struct {
	Phone string `json:"phone" binding:"required"`
}

type EditPasswordPayload struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
