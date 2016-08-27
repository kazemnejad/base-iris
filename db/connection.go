package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"yonje/baseframework/config"
	"fmt"
)

var Connection *dbr.Connection = nil

func OpenGlobalConnection() (err error) {
	if Connection != nil {
		err = nil
		return
	}

	dbConf := config.GetDatabaseConfig()
	dsn := dbConf.UserName + ":" + dbConf.Password +
		"@tcp(" + dbConf.Addr + ")/" +
		dbConf.DbName

	fmt.Println(dsn)

	Connection, err = dbr.Open(dbConf.Driver, dsn, nil)

	return
}