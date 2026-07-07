package routes

import (
	"github.com/gin-gonic/gin"

	"StudentStats-backend-go/server/handlers"
)

func RegisterExamRoutes(rg *gin.RouterGroup, examHandler *handlers.ExamHandler) {
	exams := rg.Group("/exams")
	{
		exams.GET("/", examHandler.GetAll)
		exams.GET("/:id", examHandler.GetByID)
		exams.POST("/", examHandler.Create)
	}
}
