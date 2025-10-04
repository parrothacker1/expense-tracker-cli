package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) (err error) {
	if DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{}); err != nil {
		return
	}
	if err = DB.AutoMigrate(&Expense{}); err != nil {
		return
	}
	return
}
