package routes

import (
	"StudentStats-backend-go/server/ws"
	"StudentStats-backend-go/server/handlers"
	"StudentStats-backend-go/server/repositories"
	"StudentStats-backend-go/server/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, websocketHub *ws.Hub, db *gorm.DB) {
	
	router.GET("/ws", func(c *gin.Context) {
		ws.ServeWS(websocketHub, c)
	})	
	
	api := router.Group("/api")
	{
		repos := repositories.NewContainer(db)
		services := services.NewContainer(repos, websocketHub)
		handlers := handlers.NewContainer(services)

		RegisterStudentRoutes(api, handlers.Student)
		RegisterExamRoutes(api, handlers.Exam)
		RegisterPaymentRoutes(api, handlers.Payment)

		RegisterWSRoute(api, handlers.Scan)
	}
}
