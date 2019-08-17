package config

import (
	"github.com/BurntSushi/toml"
	"runtime"
)

type GLogStashToml struct {
	Version string
	Elastic string
}

var gLogStashToml *GLogStashToml

func init()  {
	var tomlPath string
	if runtime.GOOS == `windows` {
		tomlPath = "e:/glogstash/glogstash.toml"
	} else {
		tomlPath = "/config/glogstash.toml"
	}
	_, err := toml.DecodeFile(tomlPath, &gLogStashToml)
	if err != nil{
		panic(err)
	}
}

func GetElasticConfig() string {
	if nil==gLogStashToml {
		panic("get config error")
	}
	return gLogStashToml.Elastic
}