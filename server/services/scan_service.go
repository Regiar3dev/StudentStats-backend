package services

import (
	"errors"
	"StudentStats-backend-go/server/repositories"
	"StudentStats-backend-go/server/ws"
)

type ScanService struct {
	studentRepo	*repositories.StudentRepository
	hub	*ws.Hub
}

func NewScanService(repo *repositories.StudentRepository, hub *ws.Hub) *ScanService {
	return &ScanService{
		studentRepo: repo,
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

	sent := s.hub.SendToDevice(deviceID, screenPayload)
	if !sent {
		return errors.New("La pantalla del dispositivo no se encuentra en línea")
	}

	return nil
}
