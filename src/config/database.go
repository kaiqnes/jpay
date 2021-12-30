package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
