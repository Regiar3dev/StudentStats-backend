package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"type:varchar(50);unique;not null" json:"username"`
	Password	string	`gorm:"type:varchar(255);not null" json:"-"`
	Role	string	`gorm:"type:varchar(20);default:'admin'" json:"role"`
}
