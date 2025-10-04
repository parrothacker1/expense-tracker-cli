package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Amount   float64   `gorm:"not null"`
	Date     time.Time `gorm:"not null"`
	Category string
	Note     string
}
