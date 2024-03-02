package config

type Config struct {
	Server Server
	Mysql  map[string]MysqlConf
}
type Server struct {
	Addr string `yaml:"addr"`
}

var Conf Config

func InitConf(env string) {
	LoadConf(env+".yaml", &Conf)
}
