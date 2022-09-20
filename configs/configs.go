package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

// AppConf app配置
type AppConf struct {
	Db   map[string]mongodb
	Port string
}
type mongodb struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func getEnv() string {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	return env
}

var (
	once      sync.Once
	AppConfig *AppConf
)

func Initialize() {
	once.Do(func() {
		env := getEnv()
		fmt.Println("env:" + env)

		filePath, err := filepath.Abs("configs/" + env + ".conf.toml")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &AppConfig); err != nil {
			fmt.Println(err)
		}

		// 初始化其它配置数据
		AppConfig.Port = os.Getenv("APP_PORT")
		fmt.Printf("config: %+v", *AppConfig)
	})
}
