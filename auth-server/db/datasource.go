package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=secret dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	open, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return open, nil
}
