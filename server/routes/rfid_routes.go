package routes

import (
	"StudentStats-backend-go/server/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterWSRoute(rg *gin.RouterGroup, handler *handlers.ScanHandler) {
	rfid := rg.Group("/scan")
	{
		rfid.POST("", handler.ProcessCard)
	}
}
