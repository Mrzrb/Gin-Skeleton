package config

import "time"

type MysqlConf struct {
	Service           string        `yaml:"service"`
	DataBase          string        `yaml:"database"`
	Addr              string        `yaml:"addr"`
	User              string        `yaml:"user"`
	Password          string        `yaml:"password"`
	Charset           string        `yaml:"charset"`
	InterpolateParams string        `yaml:"interpolateParams"`
	MaxIdleConns      int           `yaml:"maxidleconns"`
	MaxOpenConns      int           `yaml:"maxopenconns"`
	ConnMaxIdlTime    time.Duration `yaml:"maxIdleTime"`
	ConnMaxLifeTime   time.Duration `yaml:"connMaxLifeTime"`
	ConnTimeOut       time.Duration `yaml:"connTimeOut"`
	WriteTimeOut      time.Duration `yaml:"writeTimeOut"`
	ReadTimeOut       time.Duration `yaml:"readTimeOut"`

	// sql 字段最大长度
	MaxSqlLen int `yaml:"maxSqlLen"`

	// db metrics switch
	CollectMetrics bool `yaml:"collectMetrics"`
}

func (conf *MysqlConf) CheckConf() {
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 50
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 50
	}
	if conf.ConnMaxIdlTime == 0 {
		conf.ConnMaxIdlTime = 5 * time.Minute
	}
	if conf.ConnMaxLifeTime == 0 {
		conf.ConnMaxLifeTime = 30 * time.Minute
	}
	if conf.ConnTimeOut == 0 {
		conf.ConnTimeOut = 3 * time.Second
	}
	if conf.WriteTimeOut == 0 {
		conf.WriteTimeOut = 1200 * time.Millisecond
	}
	if conf.ReadTimeOut == 0 {
		conf.ReadTimeOut = 1200 * time.Millisecond
	}
	if conf.InterpolateParams == "" {
		// 使用string，方便后续修改默认值
		conf.InterpolateParams = "false"
	}
	if conf.MaxSqlLen == 0 {
		// 日志中sql字段长度：
		// 如果不指定使用默认2048；如果<0表示不展示sql语句；否则使用用户指定的长度，过长会被截断
		conf.MaxSqlLen = 2048
	}
}
