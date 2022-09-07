package env

import "os"

// App Env
var APP_ENV string
var APP_PORT string

// DB Envs
var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_SSL_MODE string
var DB_TIMEZONE string

func InitEnv() {
	// App Env
	APP_ENV = os.Getenv("APP_ENV")
	APP_PORT = os.Getenv("APP_PORT")

	// DB Env
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_SSL_MODE = os.Getenv("DB_SSL_MODE")
	DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
}
