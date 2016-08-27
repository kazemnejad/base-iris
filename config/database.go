package config

import "github.com/gocraft/dbr"

var DbConnection *dbr.Connection = nil

type Database struct {
	Driver   string
	Addr     string
	DbName   string
	UserName string
	Password string
}

func GetDatabaseConfig() Database {
	return Database{
		Driver: "mysql",
		Addr: "127.0.0.1:5432",
		DbName:"base_framework",
		UserName:"root",
		Password:"12345",
	}
}
