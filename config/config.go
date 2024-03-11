package config

import (
	"fmt"
)

type Config struct {
	Server       Server               `yaml:"server"`
	Mysql        map[string]MysqlConf `yaml:"mysql"`
	EmailNoreply EmailServer          `yaml:"emailNoreply"`
	Cache        Redis                `yaml:"cache"`
}
type Server struct {
	Addr string `yaml:"addr"`
}

type EmailServer struct {
	EmailServer   string `yaml:"emailServer"`
	EmailUser     string `yaml:"emailUser"`
	EmailPassword string `yaml:"emailPassword"`
	EmailPort     int    `yaml:"emailPort"`
}

type Redis struct {
	Server string `yaml:"server"`
}

var Conf Config

func InitConf(env string) {
	LoadConf(env+".yaml", &Conf)
	fmt.Printf("%+v", Conf)
}
