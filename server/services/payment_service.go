package services

import (
	"StudentStats-backend-go/server/domain"
	"StudentStats-backend-go/server/repositories"
)

type PaymentService struct {
	repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{
		repo: repo,
	}
}

func (s *PaymentService) GetAll(page, limit int) (*domain.Pagination, error) {
	return s.repo.FindAll(page, limit)
}

func (s *PaymentService) GetByID(id uint) (*domain.Payment, error) {
	return s.repo.FindByID(id)
}

func (s *PaymentService) Create(payment *domain.Payment) error {
	return s.repo.Create(payment)
}
