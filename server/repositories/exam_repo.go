package repositories

import (
	"StudentStats-backend-go/server/config"
	"StudentStats-backend-go/server/domain"
	"math"

	"gorm.io/gorm"
)

type ExamRepository struct {
	DB *gorm.DB
}

func NewExamRepository() *ExamRepository {
	return &ExamRepository{
		DB: config.DB,
	}
}

func (r *ExamRepository) FindAll(page, limit int) (*domain.Pagination, error) {
	var exams []domain.Exam
	var totalRecords int64

	if err := r.DB.Model(&domain.Exam{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	if err := r.DB.Preload("Subject").Limit(limit).Offset(offset).Find(&exams).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return &domain.Pagination{
		TotalRecords: totalRecords,
		TotalPages: totalPages,
		CurrentPage: page,
		Limit: limit,
		Data: exams,
	}, nil
}

func (r *ExamRepository) FindByID(id uint) (*domain.Exam, error) {
	var exam domain.Exam
	result := r.DB.Preload("Subject").First(&exam, id)
	return &exam, result.Error
}

func (r *ExamRepository) Create(exam *domain.Exam) error {
	return r.DB.Create(exam).Error
}

func (r *ExamRepository) Delete(id uint) error {
	exam, err := r.FindByID(id)
	if err != nil {
		return err
	}
		
	return r.DB.Delete(exam).Error
}
