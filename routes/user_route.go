package routes

import (
	"chap3-challenge2/controller"
	"chap3-challenge2/repository"
	"chap3-challenge2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	auth := router.Group("/auth")
	{
		register := auth.Group("/register")
		{
			register.POST("/user", userController.CreateUser)
			register.POST("/admin", userController.CreateAdmin)
		}
		auth.POST("/login", userController.Login)
	}

}
