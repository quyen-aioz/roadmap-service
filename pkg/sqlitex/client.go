package sqlitex

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB

func InitDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	instance = db
	return db, nil
}

func Get() (*gorm.DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	return instance, nil
}
