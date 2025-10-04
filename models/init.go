package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(dsn string) (err error) {
	if DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Silent),
	}); err != nil {
		return
	}
	if err = DB.AutoMigrate(&Expense{}); err != nil {
		return
	}
	return
}
