package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN() string {
	tidbHost := getEnvWithDefault("TIDB_HOST", "127.0.0.1")
	tidbPort := getEnvWithDefault("TIDB_PORT", "4000")
	tidbUser := getEnvWithDefault("TIDB_USER", "root")
	tidbPassword := getEnvWithDefault("TIDB_PASSWORD", "")
	tidbDBName := getEnvWithDefault("TIDB_DB_NAME", "test")
	useSSL := getEnvWithDefault("USE_SSL", "true")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		tidbUser, tidbPassword, tidbHost, tidbPort, tidbDBName, useSSL)
}

func getEnvWithDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
