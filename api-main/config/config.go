package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

type Conf struct {
	Mongo Mongo `yaml:"mongo"`
	Gin   Gin   `yaml:"gin"`
}

type Mongo struct {
	Uri string `yaml:"uri"`
}

type Gin struct {
	Address string `yaml:"address"`
}


var Config Conf


// 初始化Config, 读取config.yaml文件
func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("[config init error] ioutil.ReadFile 配置文件读取失败 " + err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println("[config init error] yaml.Unmarshal 配置文件解析失败 " + err.Error())
		return
	}
	fmt.Println("[Success] 配置文件读取成功！！！")
}

