package utils

import (
	"fmt"
	"log"

	"daijoubuteam.xyz/se214-library-management/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(config config.Config) *sqlx.DB {

	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", config.DatabaseConfig.User, config.DatabaseConfig.Password, config.DatabaseConfig.Host, config.DatabaseConfig.Port, config.DatabaseConfig.Name)
	db, err := sqlx.Connect(config.DatabaseConfig.Driver, connStr)
	if err != nil {
		log.Panic(err)
	}
	return db
}
