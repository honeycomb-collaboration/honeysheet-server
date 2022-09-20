package web

import (
	"fmt"
	"spreadsheet-server/configs"
	"spreadsheet-server/internal/model"
	"spreadsheet-server/web/route"
)

func init() {
	fmt.Println("System init start!")
	// 初始化配置信息
	configs.Initialize()
	model.ConnectDB()
	// todo 一些初始化工作，读配置文件，链接数据库
}

func Run() {
	port := configs.AppConfig.Port

	if len(port) == 0 {
		port = ":80"
	} else {
		port = ":" + port
	}

	r := route.RegisterRoute()
	err := r.Run(port)
	if err != nil {
		panic(err)
		return
	}
}
