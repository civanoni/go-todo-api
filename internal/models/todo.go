package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Completed   bool   `gorm:"default:false" json:"completed"`
}
