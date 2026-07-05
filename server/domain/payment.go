package domain

import "time"

type Payment struct {
	ID	uint	`gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID	uint	`gorm:"not null" json:"student_id"`
	Amount	float64	`gorm:"type:decimal(10,2);not null" json:"amount"`
	DueDate	time.Time	`gorm:"type:date;not null" json:"due_date"`
	IsPaid	bool	`gorm:"default:false" json:"is_paid"`
	PaidAt	*time.Time	`json:"paid_at,omitempty"`
	Student	Student	`gorm:"foreignKey:StudentID" json:"student,omitempty"`
}
