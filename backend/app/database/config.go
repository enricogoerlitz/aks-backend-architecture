package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DatabaseConnector = DatabaseConnectorFactory{}

type DatabaseConnectorFactory struct{}

func (c *DatabaseConnectorFactory) Connect(
	writeDB *gorm.DB,
	readDB *gorm.DB,
	dbType string,
	writeCfn DatabaseConfig,
	readCfn DatabaseConfig,
) error {
	switch dbType {
	case "POSTGRESQL":
		return ConnectPostgres(writeDB, readDB, writeCfn, readCfn)
	// case "MSSQL":
	// 	return ConnectMicrosoftSQLServer(writeDB, readDB, writeCfn, readCfn)
	default:
		return fmt.Errorf("database type %s not supported", dbType)
	}
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  *string
	Timezone *string
	Driver   *string
	Params   *string
}
