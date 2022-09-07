package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID      `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"created_at" gorm:"<-:create"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FirstName       string         `json:"first_name"`
	LastName        *string        `json:"last_name"`
	Email           string         `json:"email"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	Phone           *string        `json:"phone"`
	PhoneVerifiedAt *time.Time     `json:"phone_verified_at"`
	Password        string         `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
