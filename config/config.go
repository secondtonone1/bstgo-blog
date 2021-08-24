package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type MongoCfg struct {
	User        string `toml: "user"`
	Passwd      string `toml: "passwd"`
	Host        string `toml: "host"`
	MaxPoolSize int16  `toml: "maxpoolsize"`
	MaxConIdle  string `toml:"maxconidle"`
	ConTimeOut  string `toml: "contimeout"`
	Database    string `toml: "database"`
}

type CookieCfg struct {
	Host  string `toml: "host"`
	Alive int    `toml: "alive"`
}

type TotalCfg struct {
	Mongo     MongoCfg  `toml: "mongo"`
	Cookie    CookieCfg `toml: "cookie"`
	Location_ Location  `toml:"location"`
}

type Location struct {
	TimeZone string `toml:"timezone"`
}

var TotalCfgData TotalCfg

func init() {
	if _, err := toml.DecodeFile("./config/config.toml", &TotalCfgData); err != nil {
		log.Println("decode file failed , error is ", err)
		panic("decode file failed")
	}

}
