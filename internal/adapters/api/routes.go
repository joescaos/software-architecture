package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, productHandler *ProductHandler) {
	api := router.Group("/api/v1/products")
	{
		api.POST("/", productHandler.CreateProduct)
		api.GET("/:id", productHandler.GetProductByID)
		api.PUT("/:id", productHandler.UpdateProduct)
		api.DELETE("/:id", productHandler.DeleteProduct)
		api.GET("/", productHandler.ListProducts)
	}
}
