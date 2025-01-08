package storage

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() (*gorm.DB, error) {
	dbHost := getEnvValueByKey("DB_HOST", "localhost")
	dbPort := getEnvValueByKey("DB_PORT", "5432")
	dbUser := getEnvValueByKey("DB_USER", "postgres")
	dbPass := getEnvValueByKey("DB_PASS", "postgres")
	dbName := getEnvValueByKey("DB_NAME", "tapper")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getEnvValueByKey(key, fallback string) string {
	if value := os.Getenv(key); len(value) > 0 {
		logrus.Debugf("Key: %v; Got value: %v", key, value)
		return value
	}
	return fallback
}
