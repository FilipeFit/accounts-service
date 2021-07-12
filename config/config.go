package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	AppPort            string
	Profile            string
	DatabasePort       string
	DatabaseName       string
	DatabaseUser       string
	DatabasePassword   string
	DatabaseHost       string
	CustomerServiceUrl string
	PaymentsServiceUrl string
}

var Config config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Config = config{AppPort: os.Getenv("APP_PORT"),
		Profile:            os.Getenv("PROFILE"),
		DatabaseName:       os.Getenv("DATABASE_NAME"),
		DatabasePassword:   os.Getenv("DATABASE_PASSWORD"),
		DatabasePort:       os.Getenv("DATABASE_PORT"),
		DatabaseUser:       os.Getenv("DATABASE_USER"),
		DatabaseHost:       os.Getenv("DATABASE_HOST"),
		CustomerServiceUrl: os.Getenv("CUSTOMER_SERVICE_URL"),
		PaymentsServiceUrl: os.Getenv("PAYMENT_SERVICE_URL"),
	}

	if err := GetDBConnection(); err != nil {
		log.Fatal(err)
	}
}
