package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `gorm:"varchar(100)" json:"first_name"`
	LastName string `gorm:"varchar(100)" json:"last_name"`
	Email string `gorm:"varchar(100);unique" json:"email"`
}
