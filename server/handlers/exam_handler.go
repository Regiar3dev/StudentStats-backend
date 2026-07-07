package handlers

import (
	"net/http"
	"strconv"

	"StudentStats-backend-go/server/domain"
	"StudentStats-backend-go/server/services"

	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	service *services.ExamService
}

func NewExamHandler() *ExamHandler {
	return &ExamHandler{
		service: services.NewExamService(),
	}
}

func (h *ExamHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 { page = 1 }
	if limit < 1 { limit = 10 }

	exams, err := h.service.GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

func (h *ExamHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	exam, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exam)
}

func (h *ExamHandler) Create(c *gin.Context) {
	var exam domain.Exam
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Create(&exam); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exam)
}
