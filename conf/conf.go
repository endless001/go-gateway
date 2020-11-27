package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	confPath string
	// Conf config
	Conf = &Config{}
)

type Config struct {
	Server   *Server
	DataBase *DataBase
	Redis    *Redis
	Log      *Log
}

type Server struct {
	Address        string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
	AllowIp        []string
}

type DataBase struct {
	DriverName     string
	DataSourceName string
}

type Redis struct {
	ProxyList    []string
	Password     string
	Db           int
	ConnTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

type Log struct {
	Level string
}

func Load(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	c := Config{}
	yaml.Unmarshal(data, &c)
	Conf = &c
	return
}
