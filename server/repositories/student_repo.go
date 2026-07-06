package repositories

import (
	"StudentStats-backend-go/server/config"
	"StudentStats-backend-go/server/domain"
	"errors"
	"math"

	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

func (r *StudentRepository) FindAll(page, limit int) (*domain.Pagination, error) {
	var students []domain.Student
	var totalRecords int64
	
	if err := r.DB.Model(&domain.Student{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	if err := r.DB.Limit(limit).Offset(offset).Find(&students).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))
	
	return &domain.Pagination{
		TotalRecords: totalRecords,
		TotalPages: totalPages,
		CurrentPage: page,
		Limit: limit,
		Data: students,
	}, nil
}

func (r *StudentRepository) FindByID(id uint) (*domain.Student, error) {
	var student domain.Student
	result := r.DB.First(&student, id)
	return &student, result.Error
}

func (r *StudentRepository) Create(student *domain.Student) error {
	return r.DB.Create(student).Error
}

func (r *StudentRepository) FindByRFID(rfidUID string) (*domain.Student, error) {
	var credential domain.Credential

	err := r.DB.Preload("Student").
		Where("rfid_uid = ? AND is_active = ?", rfidUID, true).
		First(&credential).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Credencial no encontrada o inactiva")
		}
		return nil, err
	}

	return &credential.Student, nil
}
