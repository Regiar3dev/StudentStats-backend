package handlers

import (
	"StudentStats-backend-go/server/services"
) 

type HandlerContainer struct {
	Student *StudentHandler
	Payment  *PaymentHandler
	Exam *ExamHandler
	Scan *ScanHandler
}

func NewContainer(services *services.ServiceContainer) *HandlerContainer {
	return  &HandlerContainer{
		Student: NewStudentHandler(services.Student),
		Payment: NewPaymentHandler(services.Payment),
		Exam: NewExamHandler(services.Exam),
		Scan: NewScanHandler(services.Scan),
	}
}

