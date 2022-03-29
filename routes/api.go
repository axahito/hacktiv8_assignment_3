package routes

import (
	"hacktiv8_assignment_3/handlers"

	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	router := gin.Default()

	router.GET("/order", handlers.IndexOrder)
	router.GET("/order/:order", handlers.ShowOrder)
	router.POST("/order", handlers.CreateOrder)
	router.PUT("/order/:order", handlers.UpdateOrder)
	router.DELETE("/order/:order", handlers.DeleteOrder)

	return router
}
