package database

import (
	"explore-aks-backend-app-air/constants"

	"gorm.io/gorm"
)

var DBWrite *gorm.DB // repository
var DBRead *gorm.DB  // repository

func ConnectDB() {
	dbType := constants.ENV_DB_TYPE

	sslMode := "disable"
	writeCfn := DatabaseConfig{
		Host:     constants.ENV_DB_WRITE_HOST,
		Port:     constants.ENV_DB_WRITE_PORT,
		Username: constants.ENV_DB_WRITE_USER,
		Password: constants.ENV_DB_WRITE_PASSWORD,
		Database: constants.ENV_DB_NAME,
		SSLMode:  &sslMode,
		Timezone: nil,
		Driver:   nil,
		Params:   nil,
	}

	readCfn := DatabaseConfig{
		Host:     constants.ENV_DB_READ_HOST,
		Port:     constants.ENV_DB_READ_PORT,
		Username: constants.ENV_DB_READ_USER,
		Password: constants.ENV_DB_READ_PASSWORD,
		Database: constants.ENV_DB_NAME,
		SSLMode:  &sslMode,
		Timezone: nil,
		Driver:   nil,
		Params:   nil,
	}

	DatabaseConnector.Connect(DBWrite, DBRead, dbType, writeCfn, readCfn)
}
