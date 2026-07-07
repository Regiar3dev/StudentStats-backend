package services

import (
	"StudentStats-backend-go/server/domain"
	"StudentStats-backend-go/server/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{
		repo: repo,
	}
}

func (s  *StudentService) GetAll(page, limit int) (*domain.Pagination, error) {
	return  s.repo.FindAll(page, limit)
}

func (s *StudentService) GetByID(id uint) (*domain.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) Create(student *domain.Student) error {
	return s.repo.Create(student)
}
