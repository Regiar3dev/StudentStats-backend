package domain

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	FirstName	string	`gorm:"type:varchar(100);not null" json:"first_name"`
	LastName string	`gorm:"type:varchar(100);not null" json:"last_name"`
	Email	string	`gorm:"type:varchar(150);not null" json:"email"`
}
