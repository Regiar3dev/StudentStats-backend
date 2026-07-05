package handlers

import (
	"net/http"
	"StudentStats-backend-go/server/services"
	"github.com/gin-gonic/gin"
)

type ScanInput struct {
	RFIDUID	string	`json:"rfid_uid" binding:"required"`
	DeviceID string	`json:"device_id" binding:"required"`
}

type ScanHandler struct {
	scanService *services.ScanService
}

func NewScanHandler(service *services.ScanService) *ScanHandler {
	return &ScanHandler{
		scanService: service,
	}
}

func (h *ScanHandler) ProcessCard(c *gin.Context) {
	var input ScanInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Faltan parámetros obligatorios (rfid_uid, device_id)"})
		return
	}

	err := h.scanService.ExecuteScan(input.RFIDUID, input.DeviceID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"Lectura procesada y enviada a la pantalla"})
}
