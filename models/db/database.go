package db

import (
	"fmt"

	"github.com/Tiburso/GoManager/common"
	"github.com/Tiburso/GoManager/models/company"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: check how I can make this not a global variable
var DB *gorm.DB
var err error

func ConnectDatabase() error {
	host := common.GetEnvWithDefault("DB_HOST", "localhost")
	user := common.GetEnvWithDefault("DB_USER", "postgres")
	password := common.GetEnvWithDefault("DB_PASSWORD", "postgres")
	dbname := common.GetEnvWithDefault("DB_NAME", "gomanager")
	port := common.GetEnvWithDefault("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	return AutoMigrate()
}

func AutoMigrate() error {
	return DB.AutoMigrate(&company.Company{}, &company.Application{})
}
