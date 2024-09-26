package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
