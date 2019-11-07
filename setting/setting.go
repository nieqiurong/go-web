package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Application = &Config{}

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("conf/application.yml")
	if err != nil {
		log.Fatal("load application.yml fail !", err)
	}
	_ = yaml.Unmarshal(yamlFile, &Application)
}
