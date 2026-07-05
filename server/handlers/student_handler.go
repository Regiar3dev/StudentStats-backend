package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"StudentStats-backend-go/server/domain"
	"StudentStats-backend-go/server/services"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler() *StudentHandler {
	return &StudentHandler{
		service: services.NewStudentService(),
	}
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 { page = 1 }
	if limit < 1 { limit = 10 }

	students, err := h.service.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	student, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	var student domain.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}
