package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

var CONFIG = Config{}

func init() {
	pwd, _ := os.Getwd()
	content, err := ioutil.ReadFile(pwd + "/config/config.yml")
	if err != nil {
		log.Fatalf("解析config读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &CONFIG) != nil {
		log.Fatalf("解析config出错: %v", err)
	}
}

type Config struct {
	Setting struct {
		App struct {
			Name string `yaml:"name"`
			Key  string `yaml:"key"`
		}
		Database struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Database string `yaml:"database"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}
		Redis struct {
			Host        string `yaml:"host"`
			Port        string `yaml:"port"`
			MaxIdle     int    `yaml:"maxidle"`
			Active      int    `yaml:"active"`
			IdleTimeout int    `yaml:"timeout"`
			Password    string `yaml:"password"`
			Database    int    `yaml:"database"`
		}
	}
}
