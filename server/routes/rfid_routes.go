package routes

import (
	"StudentStats-backend-go/server/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterWSRoute(rg *gin.RouterGroup, scanHandler *handlers.ScanHandler) {
	rfid := rg.Group("/scan")
	{
		rfid.POST("", scanHandler.ProcessCard)
	}
}
