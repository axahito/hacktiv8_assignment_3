package routes

import (
	"hacktiv8_assignment_3/handlers"

	"github.com/gin-gonic/gin"
)

func Serve() *gin.Engine {
	router := gin.Default()

	router.GET("/customer", handlers.IndexCustomer)
	router.GET("/customer/:customer", handlers.ShowCustomer)
	router.POST("/customer", handlers.CreateCustomer)
	router.PUT("/customer/:customer", handlers.UpdateCustomer)
	router.DELETE("/customer/:customer", handlers.DeleteCustomer)

	return router
}
