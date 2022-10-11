package router

import (
	"assignment-2/controllers"
	"assignment-2/databases"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	db := databases.ConnectDB()

	orderController := controllers.NewOrderController(db)

	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", orderController.GetOrders)
		orderGroup.POST("/", orderController.CreateOrder)
		orderGroup.PUT("/:orderId", orderController.UpdatedOrder)
		orderGroup.DELETE("/:orderId", orderController.DeleteOrder)
	}
	return router
}
