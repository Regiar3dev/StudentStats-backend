package domain

import (
	"time"
)

type Credential struct {
	RFIDUID	string	`gorm:"primaryKey;type:varchar(50)" json:"rfid_uid"`
	StudentID	uint	`gorm:"not null" json:"student_id"`
	IsActive	bool	`gorm:"default:true" json:"is_active"`
	IssuedAt	time.Time	`json:"issued_at"`
	Student	Student	`gorm:"foreignKey:StudentID" json:"student,omitempty"`
}
