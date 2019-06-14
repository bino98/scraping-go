package db

import (
	"strings"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() (*gorm.DB) {
	db,err := gorm.Open("mysql", getConnectionString())
	if err != nil {
    panic(err.Error())
  }
  return db
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}

	return defaultValue
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "")
	dbname := getParamString("MYSQL_DATABASE", "scraping-go")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?parseTime=true", user, pass, protocol, host, port, dbname, dbargs)
}