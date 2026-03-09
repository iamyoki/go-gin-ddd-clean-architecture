package database

import (
	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/config"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitSqliteDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{})

	if err != nil {
		panic("failed to connect sqlite database")
	}

	return db
}
