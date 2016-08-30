package config

type Database struct {
	Driver   string
	Addr     string
	DbName   string
	UserName string
	Password string
}

func GetDatabaseConfig() Database {
	return Database{
		Driver:   "mysql",
		Addr:     "127.0.0.1:3306",
		DbName:   "base_framework",
		UserName: "root",
		Password: "12345",
	}
}
