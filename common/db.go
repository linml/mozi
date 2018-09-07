package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/config"
	"os"
)

var BaseDb *sql.DB

func init() {
	c, err := config.ReadDefault("conf.ini")
	if err != nil {
		fmt.Println("not found file")
		os.Exit(1)
	}
	user, err := c.String("database", "username")
	if err != nil {
		fmt.Println("not found file1")
		os.Exit(1)
	}
	password, err := c.String("database", "password")
	if err != nil {
		fmt.Println("not found file2")
		os.Exit(1)
	}
	host, err := c.String("database", "host")
	if err != nil {
		fmt.Println("not found file3")
		os.Exit(1)
	}
	port, err := c.Int("database", "port")
	if err != nil {
		fmt.Println("not found file4")
		os.Exit(1)
	}
	db, err := c.String("database", "db")
	if err != nil {
		fmt.Println("not found file5")
		os.Exit(1)
	}

	BaseDb, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, db))
	BaseDb.SetMaxOpenConns(5000)
	BaseDb.SetMaxIdleConns(1000)
	BaseDb.Ping()
}
