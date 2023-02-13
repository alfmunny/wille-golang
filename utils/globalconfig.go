package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/alfmunny/willego/zinterface"
)

type GlobalConfig struct {
	TcpServer zinterface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var Config *GlobalConfig

func (c *GlobalConfig) Reload() {
	data, err := ioutil.ReadFile("conf/willego.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}

}

func init() {
	Config = &GlobalConfig{
		Name:           "Willego Server",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.9",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	Config.Reload()
}
