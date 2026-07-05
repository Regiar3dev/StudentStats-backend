package domain

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name string	`gorm:"type:varchar(150);not null" json:"name"`
}
