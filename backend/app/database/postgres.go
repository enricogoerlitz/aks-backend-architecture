package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(writeDB *gorm.DB, readDB *gorm.DB, writeCfn DatabaseConfig, readCfn DatabaseConfig) error {
	writeDSN, readDSN := generatePostgresDSN(writeCfn, readCfn)

	var err error
	DBWrite, err = gorm.Open(postgres.Open(writeDSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("write DB error: %w", err)
	}

	if writeDSN == readDSN {
		DBRead = DBWrite
		return nil
	}

	DBRead, err = gorm.Open(postgres.Open(readDSN), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("read DB error: %w", err)
	}

	return nil
}

func generatePostgresDSN(writeCfg DatabaseConfig, readCfg DatabaseConfig) (string, string) {
	writeDSN := formatPostgresDSN(
		writeCfg.Host,
		writeCfg.Port,
		writeCfg.Username,
		writeCfg.Password,
		writeCfg.Database,
		*writeCfg.SSLMode,
	)

	readDSN := formatPostgresDSN(
		readCfg.Host,
		readCfg.Port,
		readCfg.Username,
		readCfg.Password,
		readCfg.Database,
		*readCfg.SSLMode,
	)

	return writeDSN, readDSN
}

func formatPostgresDSN(host string, port string, user string, password string, dbname string, sslmode string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}
