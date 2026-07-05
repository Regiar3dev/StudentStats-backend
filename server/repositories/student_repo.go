package repositories

import (
	"StudentStats-backend-go/server/config"
	"StudentStats-backend-go/server/domain"
)

type StudentRepository struct {}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (r *StudentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	result := config.DB.Find(&students)
	return students, result.Error
}

func (r *StudentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	result := config.DB.First(&student, id)
	return &student, result.Error
}

func (r *StudentRepository) Create(student *models.Student) error {
	return config.DB.Create(student).Error
}
