package main

import (
	"fmt"
	"log"

	"github.com/aboverio/user-service/env"
	"github.com/aboverio/user-service/routers"
	"github.com/aboverio/user-service/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if err := services.InitDB(); err != nil {
		log.Fatalln(err)
	}

	if env.APP_ENV != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()

	routers.Main(server)

	if err := server.Run(fmt.Sprintf(":%s", env.APP_PORT)); err != nil {
		log.Fatalln(err)
	}
}
