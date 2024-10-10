package initialize

import (
	"database/sql"

	"github.com/LeMinh0706/ChatApp/util"
)

var (
	pgd *sql.DB
)

func InitPostgres() (*sql.DB, error) {
	if pgd == nil {
		config, err := util.LoadConfig("../")
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
