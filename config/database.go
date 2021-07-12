package config

import (
	"fmt"
	"github.com/filipeFit/account-service/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func GetDBConnection() error {

	db, err := databaseConnection()
	if err != nil {
		return err
	}

	if err := executeMigrations(db); err != nil {
		return err
	}
	// Associating the database connection to the variable
	DB = db
	return nil
}

func databaseConnection() (*gorm.DB, error) {
	log.Printf("trying to connecting to database %s", Config.DatabaseHost)
	dbConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		Config.DatabaseHost,
		Config.DatabaseUser,
		Config.DatabasePassword,
		Config.DatabaseName,
		Config.DatabasePort)

	gormLogger := logger.Default.LogMode(logger.Silent)
	if Config.Profile == "debug" {
		gormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{Logger: gormLogger})

	if err != nil {
		return nil, err
	}

	if err := createDatabaseConnectionPool(db); err != nil {
		return nil, err
	}

	log.Printf("connection with %s succefull", Config.DatabaseHost)
	return db, nil
}

func createDatabaseConnectionPool(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Not possible to close SQL connection please check error: %s", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func executeMigrations(db *gorm.DB) error {
	log.Printf("executing migrations on %s", Config.DatabaseHost)
	err := db.AutoMigrate(&domain.Account{})
	if err != nil {
		return err
	}
	log.Print("migrations executed with success ")
	return nil
}
