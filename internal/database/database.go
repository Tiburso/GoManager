package database

import (
	"fmt"

	"github.com/Tiburso/GoManager/common"
	"github.com/Tiburso/GoManager/internal/application"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	host := common.GetEnvWithDefault("DB_HOST", "localhost")
	user := common.GetEnvWithDefault("DB_USER", "postgres")
	password := common.GetEnvWithDefault("DB_PASSWORD", "postgres")
	dbname := common.GetEnvWithDefault("DB_NAME", "gomanager")
	port := common.GetEnvWithDefault("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&application.Company{}, &application.Application{})

	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
