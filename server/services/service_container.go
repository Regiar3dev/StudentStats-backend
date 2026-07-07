package services

import (
	"StudentStats-backend-go/server/repositories"
	"StudentStats-backend-go/server/ws"
)

type ServiceContainer struct {
	Student	*StudentService
	Exam	*ExamService
	Payment	*PaymentService
	Scan *ScanService
}

func NewContainer(repos *repositories.RepositoryContainer, hub *ws.Hub) *ServiceContainer {
	return &ServiceContainer{
		Student: NewStudentService(repos.Students),
		Exam:	NewExamService(repos.Exams),
		Payment: NewPaymentService(repos.Payments),
		Scan: NewScanService(repos.Students, repos.Exams, repos.Payments, hub),
	}
}
