package config

import (
	"github.com/BurntSushi/toml"
	"runtime"
)

type GLogClientToml struct {
	Version string
	Server  string
}

var gLogClientToml *GLogClientToml

func init()  {
	var tomlPath string
	if runtime.GOOS == `windows` {
		tomlPath = "e:/glog/client.toml"
	} else {
		tomlPath = "/config/client.toml"
	}
	_, err := toml.DecodeFile(tomlPath, &gLogClientToml)
	if err != nil{
		panic(err)
	}
}

func GetGLogServer() string {
	if nil==gLogClientToml {
		panic("get config error")
	}
	return gLogClientToml.Server
}