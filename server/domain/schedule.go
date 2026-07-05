package domain

type ClassSchedule struct {
	ID	uint	`gorm:"primaryKey;autoIncrement" json:"id"`
	SubjectID	uint	`gorm:"not null" json:"subject_id"`
	ProfessorID	uint	`gorm:"not null" json:"professor_id"`
	DayOfWeek	int	`gorm:"not null" json:"day_of_week"`
	StartTime	string	`gorm:"type:varchar(5);not null" json:"start_time"`
	EndTime	string	`gorm:"type:varchar(5);not null" json:"end_time"`
	Room string	`gorm:"type:varchar(50);not null" json:"room"`
	Subject	Subject	`gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	Professor	Professor `gorm:"foreignKey:ProfessorID" json:"professor,omitempty"`
}

type Enrollment struct {
	StudentID	uint	`gorm:"primaryKey" json:"student_id"`
	SubjectID	uint	`gorm:"primaryKey" json:"subject_id"`
	Student	Student	`gorm:"foreignKey:StudentID" json:"-"`
	Subject	Subject	`gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
}
