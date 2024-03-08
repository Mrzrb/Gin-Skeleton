package config

type Config struct {
	Server       Server
	Mysql        map[string]MysqlConf
	EmailNoreply EmailServer `yaml:"emailNoreply"`
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

var Conf Config

func InitConf(env string) {
	LoadConf(env+".yaml", &Conf)
}
