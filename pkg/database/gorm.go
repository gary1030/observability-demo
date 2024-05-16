package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{
		Logger: newLogger(),
	})
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func newLogger() logger.Interface {
	return &gormLogger{
		config: logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	}
}
