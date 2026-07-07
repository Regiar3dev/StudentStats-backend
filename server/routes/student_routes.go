package routes

import (
	"github.com/gin-gonic/gin"

	"StudentStats-backend-go/server/handlers"
)

func RegisterStudentRoutes(rg *gin.RouterGroup, studentHandler *handlers.StudentHandler) {
	students := rg.Group("/students")
	{
		students.GET("/", studentHandler.GetAll)
		students.GET("/:id", studentHandler.GetByID)
		students.POST("/", studentHandler.Create)
	}
}
