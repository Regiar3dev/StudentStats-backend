package services

import (
	"StudentStats-backend-go/server/domain"
	"StudentStats-backend-go/server/repositories"
)

type ExamService struct {
	repo *repositories.ExamRepository
}

func NewExamService(repo *repositories.ExamRepository) *ExamService {
	return &ExamService{
		repo: repo,
	}
}

func (s *ExamService) GetAll(page, limit int) (*domain.Pagination, error) {
	return s.repo.FindAll(page, limit)
}

func (s *ExamService) GetByID(id uint) (*domain.Exam, error) {
	return s.repo.FindByID(id)
}

func (s *ExamService) Create(exam *domain.Exam) error {
	return s.repo.Create(exam)
}

func (s *ExamService) Delete(id uint) error {
	return s.repo.Delete(id)
}
