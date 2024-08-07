package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Firstname string `gorm:"not null" json:"first_name"`
	Lastname  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"not null" json:"email"`
	Tasks     []Task `json:"tasks"`
}
