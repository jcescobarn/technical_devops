package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func (mc *MySQLConfig) Connect() (*sql.DB, error) {

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mc.DBUser, mc.DBPassword, mc.DBHost, mc.DBPort, mc.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil

}
