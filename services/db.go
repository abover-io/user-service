package services

import (
	"fmt"

	"github.com/aboverio/user-service/env"
	"github.com/aboverio/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			env.DB_HOST,
			env.DB_PORT,
			env.DB_USER,
			env.DB_PASSWORD,
			env.DB_NAME,
			env.DB_SSL_MODE,
			env.DB_TIMEZONE,
		),
	}))
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.User{})

	DB = db

	return nil
}
