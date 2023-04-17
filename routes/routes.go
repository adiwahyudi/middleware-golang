package routes

import (
	"chap3-challenge2/controller"
	"chap3-challenge2/middleware"
	"chap3-challenge2/repository"
	"chap3-challenge2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(*userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(*productService)

	auth := router.Group("/auth")
	{
		register := auth.Group("/register")
		{
			register.POST("/user", userController.CreateUser)
			register.POST("/admin", userController.CreateAdmin)
		}
		auth.POST("/login", userController.Login)
	}
	productRoute := router.Group("/product", middleware.AuthMiddleware)
	{
		productRoute.POST("", productController.CreateProduct)
		productRoute.GET("", productController.GetListProducts)
		productRoute.GET("/:id", productController.GetProductByID)
		productRoute.PUT("/:id", productController.UpdateProductByID)
		productRoute.DELETE("/:id", productController.DeleteProductById)
	}

}
