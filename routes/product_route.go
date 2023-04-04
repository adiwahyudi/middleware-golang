package routes

import (
	"chap3-challenge2/controller"
	"chap3-challenge2/middleware"
	"chap3-challenge2/repository"
	"chap3-challenge2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRoute(router *gin.Engine, db *gorm.DB) {

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository)
	productController := controller.NewProductController(*productService)

	productRoute := router.Group("/product", middleware.AuthMiddleware)
	productRoute.POST("", productController.CreateProduct)
	productRoute.GET("", productController.GetListProducts)
	productRoute.GET("/:id", productController.GetProductByID)
	productRoute.PUT("/:id", productController.UpdateProductByID)
	productRoute.DELETE("/:id", productController.DeleteProductById)

}
