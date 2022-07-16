package web

import (
	"fmt"
	"spreadsheet-server/web/route"
)

func init() {
	fmt.Println("System init start!")
	// todo 一些初始化工作，读配置文件，链接数据库
}

func Run() {
	// todo 从配置文件读端口号
	port := "9898"

	r := route.RegisterRoute()
	err := r.Run(port)
	if err != nil {
		panic(err)
		return
	}
}
