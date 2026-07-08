package services

import (
	"StudentStats-backend-go/server/repositories"
	"StudentStats-backend-go/server/ws"
	"encoding/json"
	"errors"
	"fmt"
)

type ScanService struct {
	studentRepo	*repositories.StudentRepository
	examRepo	*repositories.ExamRepository
	paymentRepo	*repositories.PaymentRepository
	hub	*ws.Hub
}

func NewScanService(
	studentRepo *repositories.StudentRepository,
	examRepo *repositories.ExamRepository,
	paymentRepo *repositories.PaymentRepository,
	hub *ws.Hub,
) *ScanService {
	return &ScanService{
		studentRepo: studentRepo,
		examRepo: examRepo,
		paymentRepo: paymentRepo,
		hub: hub,
	}
}

func (s *ScanService) ExecuteScan(rfidUID string, deviceID string) error {
	student, err := s.studentRepo.FindByRFID(rfidUID)
	if err != nil {
		return err
	}

	screenPayload := map[string]interface{}{
		"status": "success",
		"student_id": student.ID,
		"student_name": student.FirstName + " " + student.LastName,
		"next_class": map[string]interface{}{
			"lugar": "Aula test",
			"hora": "--:--",
		},
		"next_exam": map[string]interface{}{
			"materia": "Sin exámenes agendados",
			"fecha": "--/--/----",
		},
		"pending_payments_amount": 0,
	}

	bytes, err := json.MarshalIndent(screenPayload, "", "    ")
	if err != nil {
		fmt.Printf("Error al formatear el payload: %v\n", err)
	} else {
		fmt.Println("\n📺 [WS PAYLOAD] Enviando a la pantalla:")
		fmt.Println(string(bytes))
	}

	sent := s.hub.SendToDevice(deviceID, screenPayload)
	if !sent {
		return errors.New("La pantalla del dispositivo no se encuentra en línea")
	}

	return nil
}
