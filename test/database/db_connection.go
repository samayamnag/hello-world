package database

import (
	"fmt"
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(dbName string) (db *sql.DB)  {
	var dbUser, dbPwd, dbHost string
	var dbPort int
	dbUser = "root"
	dbPwd = ""
	dbHost = "localhost"
	dbPort = 3306

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPwd, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type MySQLConfig struct {
	// Optional.
	Username, Password string

	// Host of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Host string

	// Port of the MySQL instance.
	//
	// If set, UnixSocket should be unset.
	Port int

	// UnixSocket is the filepath to a unix socket.
	//
	// If set, Host and Port should be unset.
	UnixSocket string
}


func (config MySQLConfig) dataStoreName(dbName string) string {
	var credentials string

	if config.Username != "" {
		credentials = config.Username

		if config.Password != "" {
			credentials += ":" + config.Password
		}
		
		credentials += credentials + "@"
	}

	if config.UnixSocket != "" {
		return fmt.Sprintf("%sunix(%s)/%s", credentials, config.UnixSocket, dbName)
	}

	return fmt.Sprintf("%stcp(%s:%d)/%s", credentials, config.Host, config.Port, dbName)
}