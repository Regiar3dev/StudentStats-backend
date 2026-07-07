package routes

import (
	"github.com/gin-gonic/gin"

	"StudentStats-backend-go/server/handlers"
)

func RegisterPaymentRoutes(rg *gin.RouterGroup) {
	paymentHandler := handlers.NewPaymentHandler()

	payments := rg.Group("/payments")
	{
		payments.GET("/", paymentHandler.GetAll)
		payments.GET("/:id", paymentHandler.GetByID)
		payments.POST("/", paymentHandler.Create)
	}
}

