package configreader

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server string
	Port string
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Url string
	Username string
	Password string
	Database string
}

//GetConfig 获取配置数据
func GetConfig() Config{
	config := Config{}
	content, err := ioutil.ReadFile("../../../config/config.yml")
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}

	fmt.Println(string(content))
	fmt.Printf("init data: %v", config)
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return config
}

