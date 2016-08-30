package config

import (
	"github.com/kataras/iris/config"
)

var SigningSecret = []byte("dY5NdgvZ80UKYGtIf3Df")

var iris config.Iris = config.Iris{
	IsDevelopment: true,
	Gzip:          false,
	ProfilePath:   "/debug/pprof",
}

func GetIrisConfig() config.Iris {
	return config.Default().MergeSingle(iris)
}
