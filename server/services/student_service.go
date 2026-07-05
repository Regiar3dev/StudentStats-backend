package services

import (
	"StudentStats-backend-go/server/domain"
	"StudentStats-backend-go/server/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService() *StudentService {
	return &StudentService{
		repo: repositories.NewStudentRepository(),
	}
}

func (s  *StudentService) GetAll() ([]models.Student, error) {
	return  s.repo.FindAll()
}

func (s *StudentService) GetByID(id uint) (*models.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) Create(student *models.Student) error {
	return s.repo.Create(student)
}
