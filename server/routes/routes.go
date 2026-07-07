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
		RegisterStudentRoutes(api)
		RegisterExamRoutes(api)

		studentRepo := repositories.NewStudentRepository()
		scanService := services.NewScanService(studentRepo, websocketHub)
		scanHandler := handlers.NewScanHandler(scanService)

		RegisterWSRoute(api, scanHandler)
	}
}
