package repositories

import "gorm.io/gorm"


type RepositoryContainer struct {
	Students *StudentRepository
	Exams *ExamRepository
	Payments *PaymentRepository
}

func NewContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		Students: NewStudentRepository(db),
		Exams:	NewExamRepository(db),
		Payments: NewPaymentRepository(db),
	}
}
