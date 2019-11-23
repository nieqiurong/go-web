package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Application = &Config{}

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"datasource"`
}

type Server struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type Db struct {
	UserName    string `yaml:"username"`
	PassWord    string `yaml:"password"`
	DbName      string `yaml:"dbName"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxOpen     int    `yaml:"maxOpen"`
	MaxLifetime string `yaml:"maxLifetime"`
}

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("conf/application.yml")
	if err != nil {
		log.Fatal("load application.yml fail !", err)
	}
	_ = yaml.Unmarshal(yamlFile, &Application)
}
