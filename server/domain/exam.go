package domain

import "time"

type Exam struct {
	ID	uint	`gorm:"primaryKey;autoIncrement" json:"id"`
	SubjectID	uint	`gorm:"not null" json:"subject_id"`
	Title	string	`gorm:"type:varchar(150);not null" json:"title"`
	Date	time.Time	`gorm:"type:datetime;not null" json:"date"`
	Room	string	`gorm:"type:varchar(50)" json:"room"`
	Subject	Subject	`gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
}
