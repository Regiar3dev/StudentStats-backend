package repositories

import (
	"StudentStats-backend-go/server/config"
	"StudentStats-backend-go/server/domain"
	"math"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		DB: config.DB,
	}
}

func (r *PaymentRepository) FindAll(page, limit int) (*domain.Pagination, error) {
	var payments []domain.Payment
	var totalRecords int64

	if err := r.DB.Model(&domain.Payment{}).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	if err := r.DB.Preload("Student").Limit(limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	return &domain.Pagination{
		TotalRecords: totalRecords,
		TotalPages: totalPages,
		CurrentPage: page,
		Limit: limit,
		Data: payments,
	}, nil
}

func (r *PaymentRepository) FindByID(id uint) (*domain.Payment, error) {
	var payment domain.Payment
	result := r.DB.Preload("Student").First(&payment, id)
	return &payment, result.Error
}

func (r *PaymentRepository) Create(payment *domain.Payment) error {
	return r.DB.Create(payment).Error
}
