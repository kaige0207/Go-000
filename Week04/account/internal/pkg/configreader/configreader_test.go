package configreader

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	fmt.Printf("File config: %v\n", config)
	fmt.Printf("port: %v\n", config.Port)
	fmt.Printf("url: %v\n", config.Mysql.Url)
	fmt.Printf("username: %v\n", config.Mysql.Username)
	fmt.Printf("password: %v\n", config.Mysql.Password)
	fmt.Printf("password: %v\n", config.Mysql.Database)

}
