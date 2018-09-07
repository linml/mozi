package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var BaseDb *sql.DB

func init() {

	user, err := Conf.String("database", "username")
	if err != nil {
		fmt.Println("Config database->username not found")
		os.Exit(1)
	}
	password, err := Conf.String("database", "password")
	if err != nil {
		fmt.Println("Config database->password not found")
		os.Exit(1)
	}
	host, err := Conf.String("database", "host")
	if err != nil {
		fmt.Println("Config database->host not found")
		os.Exit(1)
	}
	port, err := Conf.Int("database", "port")
	if err != nil {
		fmt.Println("Config database->port not found")
		os.Exit(1)
	}
	db, err := Conf.String("database", "db")
	if err != nil {
		fmt.Println("Config database->db not found")
		os.Exit(1)
	}

	BaseDb, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, db))
	BaseDb.SetMaxOpenConns(5000)
	BaseDb.SetMaxIdleConns(1000)
	BaseDb.Ping()
}
