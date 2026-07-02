package routes

import (
	"github.com/gin-gonic/gin"

	"StudentStats-backend-go/handlers"
)

func RegisterStudentRoutes(rg *gin.RouterGroup) {
	studentHandler := handlers.NewStudentHandler()

	students := rg.Group("/students")
	{
		students.GET("/", studentHandler.GetAll)
		students.GET("/:id", studentHandler.GetByID)
		students.POST("/", studentHandler.Create)
	}
}
