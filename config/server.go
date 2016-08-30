package config

import "github.com/kataras/iris/config"

var server config.Server = config.Server{
	ListeningAddr: "0.0.0.0:8080",
	Name:          "baseWebServer",
}

func GetServerConfig() config.Server {
	return config.DefaultServer().MergeSingle(server)
}
