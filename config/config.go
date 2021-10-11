package config

import (
	"log"

	"flag"

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

type RedisCfg struct {
	Host        string `toml: "host"`
	PoolSize    int    `toml: "poolsize"`
	IdleCons    int    `toml: "idlecons"`
	IdleTimeout int    `toml: "idletimeout"`
	Passwd      string `toml: "passwd"`
	DB          int    `toml: "db"`
}

type TotalCfg struct {
	Mongo     MongoCfg  `toml: "mongo"`
	Cookie    CookieCfg `toml: "cookie"`
	Location_ Location  `toml:"location"`
	Redis     RedisCfg  `toml:"redis"`
}

type Location struct {
	TimeZone string `toml:"timezone"`
}

var TotalCfgData TotalCfg

func init() {
	cfgpath := flag.String("config", "./config/config.toml", "-config ./config/config.toml")
	flag.Parse()
	if _, err := toml.DecodeFile(*cfgpath, &TotalCfgData); err != nil {
		log.Println("decode file failed , error is ", err)
		panic("decode file failed")
	}
}
