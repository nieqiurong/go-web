package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var Application = &Config{}

func (app *Config) IsDebug() bool {
	return app.App.Mode == "debug"
}

type App struct {
	Name   string `yaml:"name" json:"name"`
	Mode   string `yaml:"mode" json:"mode"`
	Author string `yaml:"author" json:"author"`
	Email  string `yaml:"email" json:"email"`
	QQ     string `yaml:"qq" json:"qq"`
}

type Config struct {
	App    App    `yaml:"app"`
	Server Server `yaml:"server"`
	Db     Db     `yaml:"datasource"`
	Redis  Redis  `yaml:"redis"`
	Jwt    Jwt    `yaml:"jwt"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Db struct {
	UserName    string        `yaml:"username"`
	PassWord    string        `yaml:"password"`
	DbName      string        `yaml:"dbName"`
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxOpen     int           `yaml:"maxOpen"`
	MaxLifetime time.Duration `yaml:"maxLifetime"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Jwt struct {
	Key  string        `yaml:"key"`
	Time time.Duration `yaml:"time"`
	Alg  string        `yaml:"alg"`
}

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("conf/application.yml")
	if err != nil {
		log.Fatal("load application.yml fail !", err)
	}
	_ = yaml.Unmarshal(yamlFile, &Application)
}
