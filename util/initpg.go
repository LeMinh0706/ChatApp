package util

import (
	"database/sql"
)

var (
	pgd *sql.DB
)

func InitPostgres() (*sql.DB, error) {
	if pgd == nil {
		config, err := LoadConfig("../")
		if err != nil {
			return nil, err
		}
		pgd, err = sql.Open(config.DBDriver, config.DBSource)
		if err != nil {
			return nil, err
		}
	}
	return pgd, nil
}
