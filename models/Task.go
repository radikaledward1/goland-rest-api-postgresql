package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	UserId      uint
	Title       string `gorm:"not null;unique_index"`
	Description string
	Done        bool `gorm:"default:false"`
}
