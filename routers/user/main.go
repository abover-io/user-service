package user

import (
	UserController "github.com/aboverio/user-service/controllers/user"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/", UserController.Create)

	r.GET("/", UserController.Me)

	r.PUT("/", UserController.Edit)

	r.PATCH("/email", UserController.EditEmail)
	r.PATCH("/phone", UserController.EditPhone)
	r.PATCH("/password", UserController.EditPassword)

	r.DELETE("/", UserController.Remove)
}
