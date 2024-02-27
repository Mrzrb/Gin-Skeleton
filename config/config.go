package config

import "app/infra"

type Config struct {
	Server Server
	Mysql  map[string]infra.MysqlConf
}
type Server struct {
	Addr string `yaml:"addr"`
}

var Conf Config

func InitConf(env string) {
	LoadConf(env+".yaml", &Conf)
}
