package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"yonje/baseframework/config"
)

var session *dbr.Session = nil

func OpenGlobalConnection() (err error) {
	if session != nil {
		err = nil
		return
	}

	dbConf := config.GetDatabaseConfig()
	dsn := dbConf.UserName + ":" + dbConf.Password +
		"@tcp(" + dbConf.Addr + ")/" +
		dbConf.DbName

	var conn *dbr.Connection
	conn, err = dbr.Open(dbConf.Driver, dsn, nil)
	session = conn.NewSession(nil)

	return
}

func Session() *dbr.Session {
	if session == nil {
		OpenGlobalConnection()
	}

	return session
}
